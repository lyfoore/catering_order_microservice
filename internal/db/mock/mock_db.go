package mock

import "github.com/lyfoore/order-service/internal/domain"

type MockDB struct {
	orders map[string]*domain.Order
}

func NewMockDB() *MockDB {
	return &MockDB{
		orders: map[string]*domain.Order{
			"1": {
				ID:     "1",
				UserID: "1",
				Items:  []string{"1", "2"},
				Status: "pending",
			},
			"2": {
				ID:     "2",
				UserID: "2",
				Items:  []string{"3", "4"},
				Status: "delivered",
			},
		},
	}
}

func (m *MockDB) GetOrder(id string) (*domain.Order, error) {
	order, exist := m.orders[id]
	if !exist {
		return nil, domain.ErrOrderNotFound
	}
	return order, nil
}
