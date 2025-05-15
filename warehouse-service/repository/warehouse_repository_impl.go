package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"warehouse-service/model"
)

type WarehouseRepositoryImpl struct{}

func NewWarehouseRepository() *WarehouseRepositoryImpl {
	return &WarehouseRepositoryImpl{}
}

func (w WarehouseRepositoryImpl) ShowWithStock(db *gorm.DB, shopId uuid.UUID, productId uuid.UUID) (WarehouseDetail, error) {
	var warehouse WarehouseDetail

	result := db.Model(&model.Warehouse{}).
		Select(`
			warehouses.id AS id,
			warehouses.shop_id AS shop_id,
			warehouses.name AS name,
			warehouses.address AS address,
			warehouses.created_at AS created_at,
			warehouses.updated_at AS updated_at,
			stocks.qty AS qty,
			stocks.locked_qty as locked_qty
		`).
		Joins("LEFT JOIN stocks on warehouses.id = stocks.warehouse_id").
		Where("warehouses.shop_id = ? AND stocks.product_id = ? AND warehouses.status = true", shopId, productId).
		First(&warehouse)

	return warehouse, result.Error
}

func (w WarehouseRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Warehouse, error) {
	var warehouse model.Warehouse

	query := db.Model(&model.Warehouse{}).
		Joins("LEFT JOIN warehouses on warehouses.id = stocks.warehouse_id")

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
