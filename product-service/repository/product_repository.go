package repository

import (
	"gorm.io/gorm"
	"product-service/model"
)

type ProductRepository interface {
	Create(db *gorm.DB, product model.Product) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Product, error)
	FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]model.Product, error)
}
