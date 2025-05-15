package repository

import (
	"gorm.io/gorm"
	"order-service/model"
)

type ShopOrderRepository interface {
	Create(db *gorm.DB, shopOrder model.ShopOrder) error
}
