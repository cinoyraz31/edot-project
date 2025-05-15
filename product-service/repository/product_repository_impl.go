package repository

import (
	"fmt"
	"gorm.io/gorm"
	"product-service/model"
	"reflect"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Count(db *gorm.DB, params map[string]interface{}) (int64, error) {
	var count int64
	query := db.Model(&model.Product{})

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

func (p ProductRepositoryImpl) Create(db *gorm.DB, product model.Product) error {
	tx := db.Begin()
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p ProductRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.Product, error) {
	var product model.Product

	query := db.Model(&model.Product{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&product)

	return product, result.Error
}

func (p ProductRepositoryImpl) FindAll(db *gorm.DB, params map[string]interface{}, options map[string]interface{}) ([]model.Product, error) {
	var products []model.Product
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
			query.Limit(value.(int))
		case "offset":
			query.Offset(value.(int))
		}
	}

	result := query.Find(&products)
	return products, result.Error
}
