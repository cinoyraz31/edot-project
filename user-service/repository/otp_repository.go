package repository

import (
	"gorm.io/gorm"
	"user-service/model"
)

type OtpRepository interface {
	Create(db *gorm.DB, otp model.OTP) error
	ActiveState(db *gorm.DB, otp model.OTP) bool
	MaxDailyAttempt(db *gorm.DB, otp model.OTP) bool
	Verified(db *gorm.DB, otp model.OTP) (model.OTP, bool)
	Validate(db *gorm.DB, otp model.OTP) (model.OTP, error)
	RemainingTimeActiveState(db *gorm.DB, otp model.OTP) (model.OTP, int)
	UsedOtpDaily(db *gorm.DB, otp model.OTP) int
}
