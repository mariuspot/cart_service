package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	pb "github.com/mariuspot/nab_cart_service/pkg/api"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Unable to connect to cart service (err: %v)", err)
	}
	defer conn.Close()

	client := pb.NewCartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// createCartResponse, err := client.CreateCart(ctx, &pb.CreateCartRequest{})
	// log.Printf("%#v, %#v", createCartResponse, err)

	cart_id := int64(15) //createCartResponse.GetCartId()

	addLineItemResponse, err := client.AddLineItem(ctx, &pb.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	log.Printf("%#v, %#v", addLineItemResponse, err)

	removeLineItemResponse, err := client.RemoveLineItem(ctx, &pb.RemoveLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 1})
	log.Printf("%#v, %#v", removeLineItemResponse, err)

	removeLineItemResponse, err = client.RemoveLineItem(ctx, &pb.RemoveLineItemRequest{CartId: cart_id, ProductId: 1})
	log.Printf("%#v, %#v", removeLineItemResponse, err)

	addLineItemResponse, err = client.AddLineItem(ctx, &pb.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	log.Printf("%#v, %#v", addLineItemResponse, err)

	emptyCartResponse, err := client.EmptyCart(ctx, &pb.EmptyCartRequest{CartId: cart_id})
	log.Printf("%#v, %#v", emptyCartResponse, err)

	addLineItemResponse, err = client.AddLineItem(ctx, &pb.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	log.Printf("%#v, %#v", addLineItemResponse, err)
	addLineItemResponse, err = client.AddLineItem(ctx, &pb.AddLineItemRequest{CartId: cart_id, ProductId: 2, Quantity: 2})
	log.Printf("%#v, %#v", addLineItemResponse, err)

	getLineItemsResponse, err := client.GetLineItems(ctx, &pb.GetLineItemsRequest{CartId: cart_id})
	for _, a := range getLineItemsResponse.GetLineItem() {
		log.Printf("%#v", a)
		log.Printf("%#v", a.GetUpdatedAt())

	}
}
