package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseStockRepository interface {
	Create(db *gorm.DB, stock model.Stock) error
	Update(db *gorm.DB, warehouse model.Stock) error
	Delete(db *gorm.DB, warehouseId uuid.UUID) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Stock, error)
	CheckStock(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID) (int64, error)
	LockQty(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID, qty int) error    // for update query
	ReleaseQty(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID, qty int) error // for update query
	PayQty(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID, qty int) error     // for update query
	Count(db *gorm.DB, params map[string]interface{}) (int64, error)
	QuantityTotal(db *gorm.DB, params map[string]interface{}) (int64, error)
}
