package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	cart_service "github.com/mariuspot/nab_cart_service/pkg/api"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	//Connect to cart service
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

	//Client for communicating with cart service
	client := cart_service.NewCartServiceClient(conn)

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Create a new cart
	createCartResponse, err := client.CreateCart(ctx, &cart_service.CreateCartRequest{})
	if err != nil {
		log.Fatalf("Failed to create cart (err: %s)", err.Error())
	}

	cart_id := createCartResponse.GetCartId()
	log.Printf("Cart created, cart_id:%d", cart_id)

	//Add a item to the cart
	_ /*addLineItemResponse*/, err = client.AddLineItem(ctx, &cart_service.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	if err != nil {
		log.Fatalf("Failed to add item to cart (err: %s)", err.Error())
	}
	log.Println("Items added to cart (3 of product 1)")

	//Remove a single item (quantity 1) product from cart
	_ /*removeLineItemResponse*/, err = client.RemoveLineItem(ctx, &cart_service.RemoveLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 1})
	if err != nil {
		log.Fatalf("Failed to remove item from cart (err: %s)", err.Error())
	}
	log.Println("Item removed from cart (1 of product 1)")

	//Remove all of a specific product from cart
	_ /*removeLineItemResponse*/, err = client.RemoveLineItem(ctx, &cart_service.RemoveLineItemRequest{CartId: cart_id, ProductId: 1})
	if err != nil {
		log.Fatalf("Failed to remove item from cart (err: %s)", err.Error())
	}
	log.Println("Item all of product from cart (all of product 1)")

	//Convert cart to order (Empty cart should return an error)
	convertCartToOrderResponse, err := client.ConvertCartToOrder(ctx, &cart_service.ConvertCartToOrderRequest{CartId: cart_id, Name: "Name", Address: "Address", Email: "myemail@nab.com", PayType: cart_service.ConvertCartToOrderRequest_CARD})
	if err == nil {
		log.Fatalln("Something went wrong, expected and error and didn't receive one")
	}
	log.Printf("Expected, error converting cart to order (err: %s)", err.Error())

	//Add 3 items of product 1 to cart
	_ /*addLineItemResponse*/, err = client.AddLineItem(ctx, &cart_service.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	if err != nil {
		log.Fatalf("Failed to add item to cart (err: %s)", err.Error())
	}
	log.Println("Items added to cart (3 of product 1)")

	//Empty cart (remove everything)
	_ /*emptyCartResponse*/, err = client.EmptyCart(ctx, &cart_service.EmptyCartRequest{CartId: cart_id})
	if err != nil {
		log.Fatalf("Failed to empty cart (err: %s)", err.Error())
	}
	log.Println("Emptied cart (removed everything)")

	//Add 3 items of product 1 to cart
	_ /*addLineItemResponse*/, err = client.AddLineItem(ctx, &cart_service.AddLineItemRequest{CartId: cart_id, ProductId: 1, Quantity: 3})
	if err != nil {
		log.Fatalf("Failed to add item to cart (err: %s)", err.Error())
	}
	log.Println("Items added to cart (3 of product 1)")

	//Add 2 items of product 2 to cart
	_ /*addLineItemResponse*/, err = client.AddLineItem(ctx, &cart_service.AddLineItemRequest{CartId: cart_id, ProductId: 2, Quantity: 2})
	if err != nil {
		log.Fatalf("Failed to add item to cart (err: %s)", err.Error())
	}
	log.Println("Items added to cart (2 of product 2)")

	//Get items in cart
	getLineItemsResponse, err := client.GetLineItems(ctx, &cart_service.GetLineItemsRequest{CartId: cart_id})
	if err != nil {
		log.Fatalf("Failed to get items in cart (err: %s)", err.Error())
	}
	for i, item := range getLineItemsResponse.GetLineItem() {
		fmt.Printf("Item %d", i+1)
		log.Printf("\tTitle: %s\n\tDescription: %s\n\tImageUrl: %s\n\tQuantity: %d\n\tPrice: %0.2f", item.GetTitle(), item.GetDescription(), item.GetImageUrl(), item.GetQuantity(), item.GetPrice())

	}

	//Convert cart to order
	convertCartToOrderResponse, err = client.ConvertCartToOrder(ctx, &cart_service.ConvertCartToOrderRequest{CartId: cart_id, Name: "Name", Address: "Address", Email: "myemail@nab.com", PayType: cart_service.ConvertCartToOrderRequest_CARD})
	if err != nil {
		log.Fatalf("Failed to convert cart to order (err: %s)", err.Error())
	}
	log.Printf("Cart converted to order, order_id: %d", convertCartToOrderResponse.GetOrderId())

}
