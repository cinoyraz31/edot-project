package repository

import (
	"gorm.io/gorm"
	"shop-service/model"
)

type ShopProductRepository interface {
	Create(db *gorm.DB, shopProduct model.ShopProduct) error
	Update(db *gorm.DB, shopProduct model.ShopProduct) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.ShopProduct, error)
}
