package models

import (
	"github.com/lib/pq"
	"time"
)

type Task struct {
	ID        uint
	UserID    uint
	Title     string
	Body      string
	Status    string
	Items     pq.Int64Array `gorm:"type:integer[]"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
