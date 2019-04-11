package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mariuspot/nab_cart_service/data"
	pb "github.com/mariuspot/nab_cart_service/pkg/api"
)

type cartServiceServer struct {
	dc *data.DatabaseConnection
}

func NewCartServiceServer() (*cartServiceServer, error) {
	dc, err := data.NewDatabaseConnection("username", "password", "localhost:3306", "narnes_and_boble")
	if err != nil {
		return nil, err
	}
	css := &cartServiceServer{dc: dc}
	return css, nil
}

func (css *cartServiceServer) Close() {
	if css.dc != nil {
		css.dc.Close()
	}
}

func (cs *cartServiceServer) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	log.Println("CreateCart")
	id, err := cs.dc.CreateCart()
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error creating cart (err: %s)", err.Error()))
	}
	return &pb.CreateCartResponse{CartId: id}, nil
}

func (cs *cartServiceServer) AddLineItem(ctx context.Context, req *pb.AddLineItemRequest) (*pb.AddLineItemResponse, error) {
	log.Println("AddLineItem")
	err := cs.dc.AddLineItem(req.GetCartId(), req.GetProductId(), req.GetQuantity())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error adding line item to cart (err: %s)", err.Error()))
	}
	return &pb.AddLineItemResponse{}, nil
}

func (cs *cartServiceServer) RemoveLineItem(ctx context.Context, req *pb.RemoveLineItemRequest) (*pb.RemoveLineItemResponse, error) {
	log.Println("RemoveLineItem")
	err := cs.dc.RemoveLineItem(req.GetCartId(), req.GetProductId(), req.GetQuantity())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error removing line item from cart (err: %s)", err.Error()))
	}
	return &pb.RemoveLineItemResponse{}, nil
}

func (cs *cartServiceServer) EmptyCart(ctx context.Context, req *pb.EmptyCartRequest) (*pb.EmptyCartResponse, error) {
	log.Println("EmptyCart")
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCart not implemented")
}

func (cs *cartServiceServer) GetLineItems(ctx context.Context, req *pb.GetLineItemsRequest) (*pb.GetLineItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLineItems not implemented")
}
