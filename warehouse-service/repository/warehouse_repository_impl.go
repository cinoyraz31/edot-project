package repository

import (
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseRepositoryImpl struct{}

func NewWarehouseRepository() *WarehouseRepositoryImpl {
	return &WarehouseRepositoryImpl{}
}

func (w WarehouseRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Warehouse, error) {
	var warehouse model.Warehouse

	query := db.Model(&model.Warehouse{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&warehouse)

	return warehouse, result.Error
}

func (w WarehouseRepositoryImpl) Create(db *gorm.DB, warehouse model.Warehouse) error {
	tx := db.Begin()
	err := tx.Create(&warehouse).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WarehouseRepositoryImpl) Update(db *gorm.DB, warehouse model.Warehouse) error {
	tx := db.Begin()
	err := tx.Save(&warehouse).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
