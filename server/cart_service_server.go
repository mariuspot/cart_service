package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mariuspot/nab_cart_service/data"
	pb "github.com/mariuspot/nab_cart_service/pkg/api"
)

var (
	requestsReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "cart_service_requests",
		Help: "The total number of requests",
	})
	requestSummary = promauto.NewSummary(prometheus.SummaryOpts{
		Name: "cart_service_request_time_summary",
		Help: "The request times summary",
	})
)

type cartServiceServer struct {
	dc *data.DatabaseConnection
}

func NewCartServiceServer(db_username, db_password, db_ip string, dp_port int, db_name string) (*cartServiceServer, error) {
	dc, err := data.NewDatabaseConnection(db_username, db_password, db_ip, dp_port, db_name)
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
	start := time.Now()
	requestsReceived.Inc()
	defer requestSummary.Observe(time.Since(start).Seconds())

	log.Println("CreateCart")
	id, err := cs.dc.CreateCart()
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error creating cart (err: %s)", err.Error()))
	}
	log.Printf("cart_id:%d", id)
	return &pb.CreateCartResponse{CartId: id}, nil
}

func (cs *cartServiceServer) AddLineItem(ctx context.Context, req *pb.AddLineItemRequest) (*pb.AddLineItemResponse, error) {
	start := time.Now()
	requestsReceived.Inc()
	defer requestSummary.Observe(time.Since(start).Seconds())

	log.Printf("AddLineItem cart_id:%d prod_id:%d quantity:%d", req.GetCartId(), req.GetProductId(), req.GetQuantity())
	err := cs.dc.AddLineItem(req.GetCartId(), req.GetProductId(), req.GetQuantity())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error adding line item to cart (err: %s)", err.Error()))
	}
	return &pb.AddLineItemResponse{}, nil
}

func (cs *cartServiceServer) RemoveLineItem(ctx context.Context, req *pb.RemoveLineItemRequest) (*pb.RemoveLineItemResponse, error) {
	start := time.Now()
	requestsReceived.Inc()
	defer requestSummary.Observe(time.Since(start).Seconds())

	log.Printf("RemoveLineItem cart_id:%d prod_id:%d quantity:%d", req.GetCartId(), req.GetProductId(), req.GetQuantity())
	err := cs.dc.RemoveLineItem(req.GetCartId(), req.GetProductId(), req.GetQuantity())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error removing line item from cart (err: %s)", err.Error()))
	}
	return &pb.RemoveLineItemResponse{}, nil
}

func (cs *cartServiceServer) EmptyCart(ctx context.Context, req *pb.EmptyCartRequest) (*pb.EmptyCartResponse, error) {
	start := time.Now()
	requestsReceived.Inc()
	defer requestSummary.Observe(time.Since(start).Seconds())

	log.Printf("EmptyCart cart_id:%d", req.GetCartId())
	err := cs.dc.EmptyCart(req.GetCartId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error emptying cart (err: %s)", err.Error()))
	}
	return &pb.EmptyCartResponse{}, nil
}

func (cs *cartServiceServer) GetLineItems(ctx context.Context, req *pb.GetLineItemsRequest) (*pb.GetLineItemsResponse, error) {
	start := time.Now()
	requestsReceived.Inc()
	defer requestSummary.Observe(time.Since(start).Seconds())

	log.Println("GetLineItems")
	lineItems, err := cs.dc.GetLineItems(req.GetCartId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error getting line items for cart (err: %s)", err.Error()))
	}
	response := &pb.GetLineItemsResponse{}

	for _, li := range lineItems {
		response.LineItem = append(response.LineItem, &pb.LineItem{Title: li.Title, Description: li.Description, ImageUrl: li.Image_url.String, Quantity: li.Quantity, Price: li.Price, UpdatedAt: &timestamp.Timestamp{Seconds: li.Updated_at.Unix()}})
		log.Printf("title:%s description:%s image_url:%s quantity:%d price:%f updated:%s", li.Title, li.Description, li.Image_url.String, li.Quantity, li.Price, li.Updated_at)
	}
	return response, nil
}
