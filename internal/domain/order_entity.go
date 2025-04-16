package domain

import "errors"

type Order struct {
	ID     uint64  `json:"id,omitempty"`
	UserID uint64  `json:"user_id,omitempty"`
	Items  []int64 `json:"items,omitempty"`
	Status string  `json:"status,omitempty"`
}

var ErrOrderNotFound = errors.New("order not found")
