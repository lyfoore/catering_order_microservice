package models

import (
	"github.com/lib/pq"
)

type Order struct {
	ID       uint64        `gorm:"primaryKey;column:id;autoIncrement"`
	UserID   uint64        `gorm:"column:user_id"`
	Items    pq.Int64Array `gorm:"column:items;type:integer[]"`
	Status   string        `gorm:"column:status"`
	Message  string        `gorm:"column:message"`
	Response string        `gorm:"column:response"`
}
