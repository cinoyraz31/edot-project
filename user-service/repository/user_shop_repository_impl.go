package repository

import (
	"gorm.io/gorm"
	"user-service/model"
)

type UserShopRepositoryImpl struct{}

func NewUserShopRepository() *UserShopRepositoryImpl {
	return &UserShopRepositoryImpl{}
}

func (u UserShopRepositoryImpl) FindByPhoneNumber(db *gorm.DB, phoneNumber string) (model.UserShop, error) {
	var userShop model.UserShop
	result := db.Where("phone_number = ?", phoneNumber).First(&userShop)
	return userShop, result.Error
}

func (u UserShopRepositoryImpl) Create(db *gorm.DB, userShop model.UserShop) error {
	tx := db.Begin()
	err := tx.Create(&userShop).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u UserShopRepositoryImpl) Update(db *gorm.DB, userShop model.UserShop) error {
	tx := db.Begin()
	err := tx.Save(&userShop).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u UserShopRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.UserShop, error) {
	var userShop model.UserShop

	query := db.Model(&model.UserShop{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&userShop)

	return userShop, result.Error
}
