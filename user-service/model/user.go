package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID `gorm:"primary_key;column:id"`
	Name        string    `gorm:"column:name;type:varchar(255)"`
	Code        string    `gorm:"column:code;type:varchar(15)"`
	DateOfBirth time.Time `gorm:"column:date_of_birth;type:date"`
	PhoneNumber string    `gorm:"column:phone_number;unique;type:varchar(20)"`
	LastLoginAt time.Time `gorm:"column:last_login_at;isnull"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
