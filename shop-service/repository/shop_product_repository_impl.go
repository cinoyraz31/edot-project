package repository

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"shop-service/model"
)

type ShopProductRepositoryImpl struct{}

func NewShopProductRepository() *ShopProductRepositoryImpl {
	return &ShopProductRepositoryImpl{}
}

func (s ShopProductRepositoryImpl) FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]ShopProductResult, error) {
	var shopProducts []ShopProductResult
	query := db.
		Select(`
			shop_products.shop_id,
			shops.name AS shop_name,
			shop_products.product_id,
			shop_products.price
		`).
		Table("shop_products").
		Joins("LEFT JOIN shops ON `shops`.`id` = `shop_products`.`shop_id`")

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
			query.Limit(value.(int))
		case "offset":
			query.Offset(value.(int))
		}
	}

	result := query.Find(&shopProducts)
	return shopProducts, result.Error
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
