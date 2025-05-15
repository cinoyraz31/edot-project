package model

import (
	"github.com/google/uuid"
	"time"
)

type OrderItem struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	ShopOrderId uuid.UUID `gorm:"column:shop_order_id;type:varchar(36)"`
	ProductId   uuid.UUID `gorm:"column:product_id;type:varchar(36)"`
	WarehouseId uuid.UUID `gorm:"column:warehouse_id;type:varchar(36)"`
	Qty         int       `gorm:"column:qty;type:int"`
	Price       float64   `gorm:"column:price;type:decimal(12,2)"`
	SubTotal    float64   `gorm:"column:sub_total;type:decimal(12,2)"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime"`
}
