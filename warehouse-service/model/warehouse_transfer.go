package model

import (
	"github.com/google/uuid"
	"time"
)

type Transfer struct {
	Id              uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	FromWarehouseId uuid.UUID `gorm:"column:from_warehouse_id;type:varchar(36)"`
	ToWarehouseId   uuid.UUID `gorm:"column:to_warehouse_id;type:varchar(36)"`
	ProductId       uuid.UUID `gorm:"column:product_id;type:varchar(36)"`
	Qty             int       `gorm:"column:qty;type:int"`
	CreatedAt       time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:datetime"`
}
