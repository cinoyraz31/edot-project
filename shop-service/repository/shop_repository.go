package repository

import (
	"gorm.io/gorm"
	"shop-service/model"
)

type ShopRepository interface {
	Create(db *gorm.DB, shop model.Shop) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.Shop, error)
}
