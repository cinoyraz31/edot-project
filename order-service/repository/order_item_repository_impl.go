package repository

import (
	"gorm.io/gorm"
	"order-service/model"
)

type OrderItemRepositoryImpl struct{}

func NewOrderItemRepository() *OrderItemRepositoryImpl {
	return &OrderItemRepositoryImpl{}
}

func (o OrderItemRepositoryImpl) Create(db *gorm.DB, orderItem model.OrderItem) error {
	tx := db.Begin()
	err := tx.Create(&orderItem).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
