package repository

import (
	"fmt"
	"gorm.io/gorm"
	"order-service/model"
	"reflect"
)

type OrderRepositoryImpl struct{}

func NewOrderRepository() *OrderRepositoryImpl {
	return &OrderRepositoryImpl{}
}

func (o OrderRepositoryImpl) Create(db *gorm.DB, order model.Order) error {
	tx := db.Begin()
	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (o OrderRepositoryImpl) Update(db *gorm.DB, order model.Order) error {
	tx := db.Begin()
	err := tx.Save(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (o OrderRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Order, error) {
	var order model.Order

	query := db.Model(&model.Order{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&order)

	return order, result.Error
}

func (o OrderRepositoryImpl) FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]model.Order, error) {
	var orders []model.Order
	query := db

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

	for key, value := range options {
		switch key {
		case "limit":
			intValue, ok := value.(int)

			if ok {
				query.Limit(intValue)
			}
		case "offset":
			intValue, ok := value.(int)
			if ok {
				query.Offset(intValue)
			}
		}
	}

	result := query.Find(&orders)
	return orders, result.Error
}

func (o OrderRepositoryImpl) Count(db *gorm.DB, params map[string]interface{}) (int64, error) {
	var count int64

	query := db

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
