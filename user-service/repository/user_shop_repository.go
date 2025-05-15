package repository

import (
	"gorm.io/gorm"
	"user-service/model"
)

type UserShopRepository interface {
	Create(db *gorm.DB, userShop model.UserShop) error
	Update(db *gorm.DB, userShop model.UserShop) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.UserShop, error)
	FindByPhoneNumber(db *gorm.DB, phoneNumber string) (model.UserShop, error)
}
