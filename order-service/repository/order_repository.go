package repository

import (
	"gorm.io/gorm"
	"order-service/model"
)

type OrderRepository interface {
	Create(db *gorm.DB, order model.Order) error
	Update(db *gorm.DB, order model.Order) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Order, error)
	FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]model.Order, error)
	Count(db *gorm.DB, params map[string]interface{}) (int64, error)
}
