package repository

import (
	"gorm.io/gorm"
	"order-service/model"
)

type OrderItemRepository interface {
	Create(db *gorm.DB, orderItem model.OrderItem) error
}
