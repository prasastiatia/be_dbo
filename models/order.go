package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        string    `json:"items"`
	Quantity     int       `json:"quantity"`
	TotalBayar   float32   `json:"total_bayar"`
}
