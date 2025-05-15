package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop-service/model"
)

type ShopProductResult struct {
	ShopID    uuid.UUID
	ShopName  string
	ProductId uuid.UUID
	Price     float64
}

type ShopProductRepository interface {
	Create(db *gorm.DB, shopProduct model.ShopProduct) error
	Update(db *gorm.DB, shopProduct model.ShopProduct) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.ShopProduct, error)
	FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]ShopProductResult, error)
}
