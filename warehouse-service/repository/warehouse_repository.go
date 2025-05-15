package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"warehouse-service/model"
)

type WarehouseDetail struct {
	Id        uuid.UUID
	ShopId    uuid.UUID
	Name      string
	Address   string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductId uuid.UUID
	Qty       int
	LockedQty int
}

type WarehouseRepository interface {
	Create(db *gorm.DB, warehouse model.Warehouse) error
	Update(db *gorm.DB, warehouse model.Warehouse) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Warehouse, error)
	ShowWithStock(db *gorm.DB, shopId uuid.UUID, productId uuid.UUID) (WarehouseDetail, error)
}
