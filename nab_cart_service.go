package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	pb "github.com/mariuspot/nab_cart_service/pkg/api"
	"github.com/mariuspot/nab_cart_service/server"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")

	db_username = flag.String("db_username", "username", "Database username")
	db_password = flag.String("db_password", "password", "Database password")
	db_name     = flag.String("db_name", "narnes_and_boble", "Database name")
	db_ip       = flag.String("db_ip", "127.0.0.1", "Database ip")
	db_port     = flag.Int("db_port", 3306, "Database port")

	metrics_ip   = flag.String("metric_ip", "127.0.0.1", "Metrics ip (prometheus)")
	metrics_port = flag.Int("metric_port", 2112, "Metrics port (prometheus)")

	ip   = flag.String("ip", "127.0.0.1", "The ip to listen on for the server")
	port = flag.Int("port", 10000, "The port to listen on for the server")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf("%s:%d", *metrics_ip, *metrics_port), nil)

	csServer, err := server.NewCartServiceServer(*db_username, *db_password, *db_ip, *db_port, *db_name)
	defer csServer.Close()

	pb.RegisterCartServiceServer(grpcServer, csServer)
	// ... // determine whether to use TLS
	log.Printf("Starting server listening on %s:%d", *ip, *port)
	grpcServer.Serve(lis)
}
