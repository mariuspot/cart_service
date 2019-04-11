package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/testdata"

	// "github.com/golang/protobuf/proto"

	"github.com/mariuspot/nab_cart_service/server"

	pb "github.com/mariuspot/nab_cart_service/pkg/api"
)

var (
	// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile   = flag.String("cert_file", "", "The TLS cert file")
	// keyFile    = flag.String("key_file", "", "The TLS key file")
	// jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	csServer, err := server.NewCartServiceServer()
	defer csServer.Close()

	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, csServer)
	// ... // determine whether to use TLS
	log.Printf("Starting server listening on 127.0.0.1:%d", *port)
	grpcServer.Serve(lis)
}
