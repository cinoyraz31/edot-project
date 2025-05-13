package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"user-service/exceptions"
	"user-service/helper"
	"user-service/model"
	"user-service/repository"
	"user-service/web/request"
	"user-service/web/response"
)

type OtpControllerImpl struct {
	DB             *gorm.DB
	OtpRepository  repository.OtpRepository
	UserRepository repository.UserRepository
}

func NewOtpController(
	DB *gorm.DB,
	otpRepository repository.OtpRepository,
	userRepository repository.UserRepository,
) *OtpControllerImpl {
	return &OtpControllerImpl{
		DB:             DB,
		OtpRepository:  otpRepository,
		UserRepository: userRepository,
	}
}

func (o OtpControllerImpl) Send(ctx *fiber.Ctx) error {
	var data request.OTPSendRequest
	var otp model.OTP
	fmt.Println(data)

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	phoneNumber, boolean := helper.ValidatePhoneNumber(data.PhoneNumber)

	if boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "No telepon wajib format +62 | 62 | 08")
	}

	otp.PhoneNumber = phoneNumber

	if boolean := o.OtpRepository.ActiveState(o.DB, otp); boolean {
		_, second := o.OtpRepository.RemainingTimeActiveState(o.DB, otp)
		return exceptions.ErrorHandlerBadRequest(ctx, fmt.Sprintf("OTP masih aktif, mohon tunggu %s detik lagi", strconv.Itoa(second)))
	}

	if boolean := o.OtpRepository.MaxDailyAttempt(o.DB, otp); boolean {
		return exceptions.ErrorHandlerBadRequest(ctx, "Maaf, anda tidak bisa mencoba kembali")
	}

	otp.Code = strconv.Itoa(helper.GenerateRandomNumber(6))
	otp.Id = uuid.New()

	if err := o.OtpRepository.Create(o.DB, otp); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, " Gagal tambah otp")
	}

	// send to OTP provider

	return ctx.Status(fiber.StatusNoContent).SendString("")
}

func (o OtpControllerImpl) Validate(ctx *fiber.Ctx) error {
	var data request.OTPValidateRequest
	var otp model.OTP

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	phoneNumber, boolean := helper.ValidatePhoneNumber(data.PhoneNumber)

	if boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "No telepon wajib format +62 | 62 | 08")
	}

	otp.PhoneNumber = phoneNumber
	otp.Code = data.Code

	if boolean := o.OtpRepository.ActiveState(o.DB, otp); boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "OTP sudah kadaluarsa")
	}

	newOtp, err := o.OtpRepository.Validate(o.DB, otp)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "OTP tidak sah")
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OTP validate success", map[string]interface{}{
		"id": newOtp.Id,
	}, nil))
}

func (o OtpControllerImpl) CountDown(ctx *fiber.Ctx) error {
	var data request.OTPSendRequest
	var otp model.OTP

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	phoneNumber, boolean := helper.ValidatePhoneNumber(data.PhoneNumber)

	if boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "No telepon wajib format +62 | 62 | 08")
	}

	otp.PhoneNumber = phoneNumber

	var wg sync.WaitGroup
	chanOtp := make(chan model.OTP, 1)
	chanSecond := make(chan int, 1)
	chanUsed := make(chan int, 1)

	wg.Add(2)

	go func() {
		defer wg.Done()
		otp, second := o.OtpRepository.RemainingTimeActiveState(o.DB, otp)
		chanOtp <- otp
		chanSecond <- second
	}()

	go func() {
		defer wg.Done()
		chanUsed <- o.OtpRepository.UsedOtpDaily(o.DB, otp)
	}()

	wg.Wait()

	close(chanOtp)
	close(chanSecond)
	close(chanUsed)

	otp = <-chanOtp
	second := <-chanSecond
	used := <-chanUsed
	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OTP Count Down", map[string]interface{}{
		"reminingTime":  second,
		"used":          used,
		"lastSendOtpAt": otp.CreatedAt,
	}, nil))
}
