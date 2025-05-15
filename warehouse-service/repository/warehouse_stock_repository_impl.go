package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"warehouse-service/model"
)

type WarehouseStockRepositoryImpl struct{}

func NewWarehouseStockRepository() *WarehouseStockRepositoryImpl {
	return &WarehouseStockRepositoryImpl{}
}

func (w WarehouseStockRepositoryImpl) Delete(db *gorm.DB, warehouseId uuid.UUID) error {
	if err := db.Where("warehouse_id = ?", warehouseId).Delete(&model.Stock{}).Error; err != nil {
		return err
	}
	return nil
}

func (w WarehouseStockRepositoryImpl) Create(db *gorm.DB, stock model.Stock) error {
	tx := db.Begin()
	err := tx.Create(&stock).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WarehouseStockRepositoryImpl) Update(db *gorm.DB, stock model.Stock) error {
	if stock.Qty < stock.LockedQty {
		return errors.New("not enough stock")
	}

	tx := db.Begin()
	err := tx.Save(&stock).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WarehouseStockRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Stock, error) {
	var stock model.Stock

	query := db.Model(&model.Stock{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&stock)

	return stock, result.Error
}

func (w WarehouseStockRepositoryImpl) CheckStock(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID) (int64, error) {
	var stock model.Stock
	result := db.Model(&model.Stock{}).
		Where("warehouse_id = ? AND product_id = ?", warehouseId, productId).
		First(&stock)

	return int64(stock.Qty - stock.LockedQty), result.Error

}

func (w WarehouseStockRepositoryImpl) LockQty(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID, qty int) error {
	tx := db.Begin()
	var stock model.Stock

	if err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("warehouse_id = ? AND product_id = ?", warehouseId, productId).
		First(&stock).Error; err != nil {
		return err
	}

	if stock.Qty-stock.LockedQty < qty {
		return errors.New("not enough stock")
	}

	stock.LockedQty += qty
	err := db.Save(&stock).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
