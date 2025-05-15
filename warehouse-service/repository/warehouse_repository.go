package repository

import (
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseRepository interface {
	Create(db *gorm.DB, warehouse model.Warehouse) error
	Update(db *gorm.DB, warehouse model.Warehouse) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Warehouse, error)
}
