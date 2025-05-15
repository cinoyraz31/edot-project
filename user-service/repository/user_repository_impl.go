package repository

import (
	"gorm.io/gorm"
	"user-service/model"
)

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) FindByPhoneNumber(db *gorm.DB, phoneNumber string) (model.User, error) {
	var user model.User
	result := db.Where("phone_number = ?", phoneNumber).First(&user)
	return user, result.Error
}

func (u UserRepositoryImpl) Create(db *gorm.DB, user model.User) error {
	tx := db.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u UserRepositoryImpl) Update(db *gorm.DB, user model.User) error {
	tx := db.Begin()
	err := tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u UserRepositoryImpl) FindBy(db *gorm.DB, params map[string]interface{}) (model.User, error) {
	var user model.User

	query := db.Model(&model.User{})

	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	result := query.First(&user)

	return user, result.Error
}
