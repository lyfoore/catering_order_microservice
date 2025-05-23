package gRPC

import (
	"context"
	"errors"
	"github.com/lyfoore/order-service/internal/domain"
	"github.com/lyfoore/order-service/internal/genproto/orders"
	"github.com/lyfoore/order-service/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type OrderServer struct {
	orders.UnimplementedOrderServiceServer
	service *service.OrderService
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.OrderResponse, error) {
	requestConverted := &domain.Order{
		UserID: req.UserId,
		Items:  req.Items,
	}
	order, err := s.service.CreateOrderService(requestConverted)
	if err != nil {
		return nil, err
	}
	return &orders.OrderResponse{
		Id:     order.ID,
		UserId: order.UserID,
		Items:  order.Items,
		Status: order.Status,
	}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.OrderResponse, error) {
	order, err := s.service.GetOrderService(req.Id)
	if err != nil {
		return nil, err
	}
	return &orders.OrderResponse{
		Id:     order.ID,
		UserId: order.UserID,
		Items:  order.Items,
		Status: order.Status,
	}, nil
}

func (s *OrderServer) UpdateOrder(ctx context.Context, req *orders.UpdateOrderRequest) (*orders.OrderResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}
	requestConverted := &domain.Order{
		ID:     req.Id,
		UserID: req.UserId,
		Items:  req.Items,
		Status: req.Status,
	}
	order, err := s.service.UpdateOrderService(requestConverted)
	if err != nil {
		return nil, err
	}
	return &orders.OrderResponse{
		Id:     order.ID,
		UserId: order.UserID,
		Items:  order.Items,
		Status: order.Status,
	}, nil
}

func (s *OrderServer) DeleteOrder(ctx context.Context, req *orders.DeleteOrderRequest) (*orders.DeleteOrderResponse, error) {
	if err := s.service.DeleteOrderService(req.Id); err != nil {
		return nil, err
	}
	return &orders.DeleteOrderResponse{
		Deleted: true,
	}, nil
}

func NewOrderServer(service *service.OrderService) *OrderServer {
	return &OrderServer{
		service: service,
	}
}

func (s *OrderServer) RunServer() {
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":9090")
	orders.RegisterOrderServiceServer(grpcServer, s)
	log.Println("gRPC server started")
	grpcServer.Serve(lis)
}
