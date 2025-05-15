package model

import (
	"github.com/google/uuid"
	"time"
)

type Stock struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	WarehouseId uuid.UUID `gorm:"column:warehouse_id;type:varchar(36)"`
	ProductId   uuid.UUID `gorm:"column:product_id;type:varchar(36)"`
	Qty         int       `gorm:"column:qty;type:int"`
	LockedQty   int       `gorm:"column:locked_qty;type:int"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime"`
}
