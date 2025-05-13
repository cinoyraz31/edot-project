package model

import (
	"github.com/google/uuid"
	"time"
)

type OTP struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:varchar(36)"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(20)"`
	Code        string    `gorm:"column:code;type:varchar(6)"`
	IsSuccess   bool      `gorm:"column:is_success;default:false"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
