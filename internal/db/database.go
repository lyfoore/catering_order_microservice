package database

import (
	"github.com/lib/pq"
	"github.com/lyfoore/order-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID       uint64        `gorm:"primaryKey;column:id;autoIncrement"`
	UserID   uint64        `gorm:"column:user_id"`
	Items    pq.Int64Array `gorm:"column:items;type:integer[]"`
	Status   string        `gorm:"column:status"`
	Message  string        `gorm:"column:message"`
	Response string        `gorm:"column:response"`
}

func (Order) TableName() string {
	return "orders"
}

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Order{}); err != nil {
		return nil, err
	}

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) CreateOrder(order *domain.Order) (*domain.Order, error) {
	orderConverted := &Order{
		UserID:   order.UserID,
		Items:    order.Items,
		Status:   order.Status,
		Message:  order.Message,
		Response: order.Response,
	}
	p.DB.Create(orderConverted)
	order.ID = orderConverted.ID
	return order, nil
}

func (p *PostgresDB) GetOrder(id uint64) (*domain.Order, error) {
	var order Order
	if err := p.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &domain.Order{
		ID:       order.ID,
		UserID:   order.UserID,
		Items:    order.Items,
		Status:   order.Status,
		Message:  order.Message,
		Response: order.Response,
	}, nil
}

func (p *PostgresDB) UpdateOrder(updatedOrder *domain.Order) error {
	order := Order{
		ID:       updatedOrder.ID,
		UserID:   updatedOrder.UserID,
		Items:    updatedOrder.Items,
		Status:   updatedOrder.Status,
		Message:  updatedOrder.Message,
		Response: updatedOrder.Response,
	}

	return p.DB.Model(&Order{}).Where("id = ?", updatedOrder.ID).Updates(order).Error
}

func (p *PostgresDB) DeleteOrder(id uint64) error {
	return p.DB.Where("id = ?", id).Delete(&Order{}).Error
}
