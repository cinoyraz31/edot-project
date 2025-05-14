package model

import (
	"github.com/google/uuid"
	"time"
)

type ShopProduct struct {
	Id        uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	ShopId    uuid.UUID `gorm:"column:shop_id;type:varchar(36)"`
	ProductId uuid.UUID `gorm:"column:product_id;type:varchar(36)"`
	Price     float64   `gorm:"column:price;type:decimal(10,2)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
