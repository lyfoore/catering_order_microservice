package sagas

import (
	"github.com/lyfoore/order-service/internal/domain"
	"github.com/lyfoore/order-service/internal/service"
	"log"
)

type OrderSaga struct {
	orderService *service.OrderService
}

func NewOrderSaga(orderService *service.OrderService) *OrderSaga {
	return &OrderSaga{
		orderService: orderService,
	}
}

func (s *OrderSaga) CreateOrder(order *domain.Order) (*domain.Order, error) {
	order.StatusAI = domain.StatusCreated
	newOrder, err := s.orderService.CreateOrderService(order)
	if err != nil {
		return nil, err
	}

	if order.Message == "" {
		return newOrder, nil
	}

	go s.processAI(newOrder.ID)

	responseOrder := *newOrder
	responseOrder.StatusAI = domain.StatusPendingAI
	if _, err = s.orderService.UpdateOrderService(&responseOrder); err != nil {
		return nil, err
	}
	log.Print(responseOrder)
	return &responseOrder, nil
}

func (s *OrderSaga) processAI(orderID uint64) {
	order, err := s.orderService.GetOrderService(orderID)
	if err != nil || order == nil {
		s.handleAIError(orderID, err)
		return
	}

	order.StatusAI = domain.StatusPendingAI
	if _, err := s.orderService.UpdateOrderService(order); err != nil {
		s.handleAIError(orderID, err)
		return
	}

	respondedOrder, err := s.orderService.GenerateAIResponseAndUpdateOrder(order)
	if err != nil {
		s.handleAIError(orderID, err)
		return
	}

	respondedOrder.StatusAI = domain.StatusCompleted
	if _, err := s.orderService.UpdateOrderService(respondedOrder); err != nil {
		s.handleAIError(orderID, err)
	}
}

func (s *OrderSaga) handleAIError(orderID uint64, err error) {
	log.Printf("Failed to process AI for order %s: %v", orderID, err)
	order, errGet := s.orderService.GetOrderService(orderID)
	if errGet != nil || order == nil {
		log.Printf("Order %s not found: %v", orderID, errGet)
		return
	}

	order.StatusAI = domain.StatusFailed
	if _, errUpdate := s.orderService.UpdateOrderService(order); errUpdate != nil {
		log.Printf("Failed to update order status: %v", errUpdate)
	}

	if errDelete := s.orderService.DeleteOrderService(orderID); errDelete != nil {
		log.Printf("Failed to delete order: %v", errDelete)
	}
}
