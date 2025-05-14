package domain

import "errors"

const (
	StatusCreated   string = "CREATED"
	StatusPendingAI string = "PENDING_AI"
	StatusCompleted string = "COMPLETED"
	StatusFailed    string = "FAILED"
)

type Order struct {
	ID       uint64  `json:"id,omitempty"`
	UserID   uint64  `json:"user_id,omitempty"`
	Items    []int64 `json:"items,omitempty"`
	Status   string  `json:"status,omitempty"`
	StatusAI string  `json:"status_ai,omitempty"`
	Message  string  `json:"message,omitempty"`
	Response string  `json:"response,omitempty"`
}

var ErrOrderNotFound = errors.New("order not found")
