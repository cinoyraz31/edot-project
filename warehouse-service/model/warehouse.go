package model

import (
	"github.com/google/uuid"
	"time"
)

type Warehouse struct {
	Id        uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	ShopId    uuid.UUID `gorm:"column:shop_id;type:varchar(36)"`
	Name      string    `gorm:"column:name;type:varchar(255)"`
	Address   string    `gorm:"column:address;type:varchar(255)"`
	Status    bool      `gorm:"column:status;type:tinyint(1)"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}
