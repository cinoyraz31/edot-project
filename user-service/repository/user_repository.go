package repository

import (
	"gorm.io/gorm"
	"user-service/model"
)

type UserRepository interface {
	Create(db *gorm.DB, user model.User) error
	Update(db *gorm.DB, user model.User) error
	FindBy(db *gorm.DB, params map[string]interface{}) (model.User, error)
}
