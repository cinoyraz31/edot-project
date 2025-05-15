package repository

import (
	"gorm.io/gorm"
	"order-service/model"
)

type ShopOrderRepositoryImpl struct{}

func NewShopOrderRepository() *ShopOrderRepositoryImpl {
	return &ShopOrderRepositoryImpl{}
}

func (s ShopOrderRepositoryImpl) Create(db *gorm.DB, shopOrder model.ShopOrder) error {
	tx := db.Begin()
	err := tx.Create(&shopOrder).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
