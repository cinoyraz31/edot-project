package repository

import (
	"gorm.io/gorm"
	"shop-service/model"
)

type ShopProductRepositoryImpl struct{}

func NewShopProductRepository() *ShopProductRepositoryImpl {
	return &ShopProductRepositoryImpl{}
}

func (s ShopProductRepositoryImpl) Create(db *gorm.DB, shopProduct model.ShopProduct) error {
	tx := db.Begin()
	err := tx.Create(&shopProduct).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s ShopProductRepositoryImpl) Update(db *gorm.DB, shopProduct model.ShopProduct) error {
	tx := db.Begin()
	err := tx.Save(&shopProduct).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s ShopProductRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.ShopProduct, error) {
	var shopProduct model.ShopProduct

	query := db.Model(&model.ShopProduct{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&shopProduct)

	return shopProduct, result.Error
}
