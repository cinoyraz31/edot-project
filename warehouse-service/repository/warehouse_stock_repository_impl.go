package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"warehouse-service/model"
)

type WarehouseStockRepositoryImpl struct{}

func NewWarehouseStockRepository() *WarehouseStockRepositoryImpl {
	return &WarehouseStockRepositoryImpl{}
}

func (w WarehouseStockRepositoryImpl) ReleaseQty(db *gorm.DB, warehouseId uuid.UUID, productId uuid.UUID, qty int) error {
	tx := db.Begin()
	var stock model.Stock

	if err := db.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("warehouse_id = ? AND product_id = ?", warehouseId, productId).
		First(&stock).Error; err != nil {
		return err
	}

	if stock.LockedQty <= 0 {
		return errors.New("locked qty must be greater than 0")
	}

	if qty > stock.LockedQty {
		return errors.New("qty must be less than or equal to locked qty")
	}

	stock.LockedQty -= qty
	err := db.Save(&stock).Error

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (w WarehouseStockRepositoryImpl) QuantityTotal(db *gorm.DB, params map[string]interface{}) (int64, error) {
	var totalQuantity int64

	query := db.Debug().
		Model(&model.Stock{}).
		Joins("LEFT JOIN warehouses on warehouses.id = stocks.warehouse_id")

	for key, value := range params {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			for k, v := range value.(map[string]interface{}) {
				query = query.Where(fmt.Sprintf("%s %s ?", key, k), v)
			}
		default:
			query = query.Where(key+" = ?", value)
		}
	}

	result := query.Select("SUM(stocks.qty-stocks.locked_qty)").Scan(&totalQuantity)
	return totalQuantity, result.Error
}

func (w WarehouseStockRepositoryImpl) Count(db *gorm.DB, params map[string]interface{}) (int64, error) {
	var count int64
	query := db.Debug().
		Model(&model.Stock{}).
		Joins("LEFT JOIN warehouses on warehouses.id = stocks.warehouse_id")

	for key, value := range params {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			for k, v := range value.(map[string]interface{}) {
				query = query.Where(fmt.Sprintf("%s %s ?", key, k), v)
			}
		default:
			query = query.Where(key+" = ?", value)
		}
	}

	result := query.Count(&count)
	return count, result.Error
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
