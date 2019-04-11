package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/mariuspot/nab_cart_service/pkg/api"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to cart service (err: %v)", err)
	}
	defer conn.Close()

	client := pb.NewCartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createCartResponse, err := client.CreateCart(ctx, &pb.CreateCartRequest{})
	log.Printf("%#v, %#v", createCartResponse, err)

	cart_id := createCartResponse.GetCartId()

	addLineItemResponse, err := client.AddLineItem(ctx, &pb.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	log.Printf("%#v, %#v", addLineItemResponse, err)

	removeLineItemResponse, err := client.RemoveLineItem(ctx, &pb.RemoveLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 1})
	log.Printf("%#v, %#v", removeLineItemResponse, err)

	removeLineItemResponse, err = client.RemoveLineItem(ctx, &pb.RemoveLineItemRequest{CartId: cart_id, ProductId: 1})
	log.Printf("%#v, %#v", removeLineItemResponse, err)

}
