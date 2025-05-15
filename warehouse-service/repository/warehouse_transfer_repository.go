package repository

import (
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseTransferRepository interface {
	Create(db *gorm.DB, warehouseTransfer model.Transfer) error
}
