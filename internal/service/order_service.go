package service

import (
	"github.com/lyfoore/order-service/internal/domain"
	"github.com/lyfoore/order-service/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) GetOrderService(id uint64) (*domain.Order, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) CreateOrderService(newOrder *domain.Order) (*domain.Order, error) {
	newOrder.Status = "created"
	order, err := s.repo.CreateOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrderService(updatedOrder *domain.Order) (*domain.Order, error) {
	existingOrder, err := s.repo.GetOrder(updatedOrder.ID)
	if err != nil {
		return nil, err
	}

	if updatedOrder.UserID != 0 {
		existingOrder.UserID = updatedOrder.UserID
	}
	if updatedOrder.Items != nil {
		existingOrder.Items = updatedOrder.Items
	}
	if updatedOrder.Status != "" {
		existingOrder.Status = updatedOrder.Status
	}

	if err := s.repo.UpdateOrder(existingOrder); err != nil {
		return nil, err
	}
	return existingOrder, nil
}

func (s *OrderService) DeleteOrderService(id uint64) error {
	return s.repo.DeleteOrder(id)
}
