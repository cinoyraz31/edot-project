package repository

import (
	"gorm.io/gorm"
	"shop-service/model"
)

type ShopRepositoryImpl struct{}

func NewShopRepository() *ShopRepositoryImpl {
	return &ShopRepositoryImpl{}
}

func (s ShopRepositoryImpl) Create(db *gorm.DB, shop model.Shop) error {
	tx := db.Begin()
	err := tx.Create(&shop).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s ShopRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Shop, error) {
	var shop model.Shop

	query := db.Model(&model.Shop{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&shop)

	return shop, result.Error
}
