package repository

import (
	"github.com/lyfoore/order-service/internal/domain"
)

type OrderRepository interface {
	GetOrder(id uint64) (*domain.Order, error)
	CreateOrder(order *domain.Order) (*domain.Order, error)
	UpdateOrder(updatedOrder *domain.Order) error
	DeleteOrder(id uint64) error
}
