package repository

import (
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseTransferRepositoryImpl struct{}

func NewWarehouseTransferRepository() *WarehouseTransferRepositoryImpl {
	return &WarehouseTransferRepositoryImpl{}
}

func (w WarehouseTransferRepositoryImpl) Create(db *gorm.DB, warehouseTransfer model.Transfer) error {
	tx := db.Begin()
	err := tx.Create(&warehouseTransfer).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
