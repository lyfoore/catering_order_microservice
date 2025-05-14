package service

import (
	"context"
	"github.com/lyfoore/order-service/internal/domain"
	"github.com/lyfoore/order-service/internal/repository"
	"github.com/openai/openai-go"
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
	if updatedOrder.StatusAI != "" {
		existingOrder.StatusAI = updatedOrder.StatusAI
	}
	if updatedOrder.Response != "" {
		existingOrder.Response = updatedOrder.Response
	}

	if err := s.repo.UpdateOrder(existingOrder); err != nil {
		return nil, err
	}
	return existingOrder, nil
}

func (s *OrderService) DeleteOrderService(id uint64) error {
	return s.repo.DeleteOrder(id)
}

func (s *OrderService) GetResponseFromAI(order *domain.Order) (*domain.Order, error) {
	client := openai.NewClient()
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(order.Message),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		return nil, err
	}

	respondedOrder := *order
	respondedOrder.Response = chatCompletion.Choices[0].Message.Content
	err = s.repo.UpdateOrder(&respondedOrder)
	if err != nil {
		return nil, err
	}

	return &respondedOrder, nil
}
