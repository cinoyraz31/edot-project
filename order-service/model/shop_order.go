package model

import (
	"github.com/google/uuid"
	"time"
)

type ShopOrder struct {
	Id        uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	OrderId   uuid.UUID `gorm:"column:order_id;type:varchar(36)"`
	ShopId    uuid.UUID `gorm:"column:shop_id;type:varchar(36)"`
	SubTotal  float64   `gorm:"column:sub_total;type:decimal(12,2)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
