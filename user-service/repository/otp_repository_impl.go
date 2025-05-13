package repository

import (
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
	"user-service/model"
)

type OTPRepositoryImpl struct{}

func NewOTPRepository() *OTPRepositoryImpl {
	return &OTPRepositoryImpl{}
}

func (otpRepository OTPRepositoryImpl) UsedOtpDaily(db *gorm.DB, otp model.OTP) int {
	var count int64
	db.Model(&otp).Where("phone_number = ? and is_success = ? and DATE(created_at) = ?", otp.PhoneNumber, false, time.Now().Format("2006-01-02")).Count(&count)
	return int(count)
}

func (otpRepository OTPRepositoryImpl) RemainingTimeActiveState(db *gorm.DB, otp model.OTP) (model.OTP, int) {
	timeSecond, _ := strconv.Atoi(os.Getenv("OTP_REMAINING_TIME"))
	remainingTime := time.Now().Add(time.Duration(-1*timeSecond) * time.Second)

	result := db.Where("phone_number = ? and is_success = ? and created_at >= ?", otp.PhoneNumber, false, remainingTime).First(&otp)

	if result.Error != nil {
		return otp, 0
	}

	var diffInSeconds int64
	db.Raw("SELECT TIMESTAMPDIFF(SECOND, MIN(?), MAX(?))", otp.CreatedAt, time.Now()).Scan(&diffInSeconds)

	remainingSecond := timeSecond - int(diffInSeconds)

	if remainingSecond > 0 {
		return otp, remainingSecond
	}
	return otp, 0
}

func (otpRepository OTPRepositoryImpl) Validate(db *gorm.DB, otp model.OTP) (model.OTP, error) {
	timeSecond, _ := strconv.Atoi(os.Getenv("OTP_REMAINING_TIME"))
	remainingTime := time.Now().Add(time.Duration(-1*timeSecond) * time.Second)

	result := db.Where("phone_number = ? and is_success = ? and created_at >= ?", otp.PhoneNumber, false, remainingTime)
	result.First(&otp)

	if result.Error != nil {
		return otp, result.Error
	}

	db.Model(&otp).Where("id = ?", otp.Id).UpdateColumn("is_success", true)
	return otp, nil
}

func (otpRepository OTPRepositoryImpl) Verified(db *gorm.DB, otp model.OTP) (model.OTP, bool) {
	timeSecond, _ := strconv.Atoi(os.Getenv("OTP_REMAINING_TIME"))
	remainingTime := time.Now().Add(time.Duration(-1*timeSecond) * time.Second)

	result := db.Where("id = ? and is_success = ? and created_at >= ?", otp.Id, true, remainingTime).First(&otp)

	if result.Error != nil {
		return otp, false
	}
	return otp, true
}

func (otpRepository OTPRepositoryImpl) MaxDailyAttempt(db *gorm.DB, otp model.OTP) bool {
	maxDaily, _ := strconv.Atoi(os.Getenv("OTP_MAX_DAILY"))
	var count int64
	db.Model(&otp).Where("phone_number = ? and DATE(created_at) = ?", otp.PhoneNumber, time.Now().Format("2006-01-02")).Count(&count)

	return count >= int64(maxDaily)
}

func (otpRepository OTPRepositoryImpl) ActiveState(db *gorm.DB, otp model.OTP) bool {
	timeSecond, _ := strconv.Atoi(os.Getenv("OTP_REMAINING_TIME"))
	var count int64

	remainingTime := time.Now().Add(time.Duration(-1*timeSecond) * time.Second)
	db.Model(&otp).Where("phone_number = ? and is_success = ? and created_at >= ?", otp.PhoneNumber, false, remainingTime).Count(&count)

	if count > 0 {
		return true
	}
	return false
}

func (otpRepository OTPRepositoryImpl) Create(db *gorm.DB, otp model.OTP) error {
	tx := db.Begin()
	err := tx.Create(&otp).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
