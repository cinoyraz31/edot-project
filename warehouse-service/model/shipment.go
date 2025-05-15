package model

import (
	"github.com/google/uuid"
	"time"
)

type Shipment struct {
	Id             uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	OrderId        uuid.UUID `gorm:"column:order_id;type:varchar(36)"`
	WarehouseId    uuid.UUID `gorm:"column:warehouse_id;type:varchar(36)"`
	Courier        string    `gorm:"column:courier;type:varchar(255)"`
	TrackingNumber string    `gorm:"column:tracking_number;type:varchar(255)"`
	Status         string    `gorm:"column:status;type:varchar(64)"` // pending / packed / sent / delivered / failed
	ShippedAt      time.Time `gorm:"column:shipped_at;type:datetime"`
	DeliveredAt    time.Time `gorm:"column:delivered_at;type:datetime"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime"`
}
