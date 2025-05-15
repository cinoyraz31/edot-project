package model

import (
	"github.com/google/uuid"
	"time"
)

type UserShop struct {
	Id          uuid.UUID `gorm:"primary_key;column:id"`
	PhoneNumber string    `gorm:"column:phone_number;unique;type:varchar(20)"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
