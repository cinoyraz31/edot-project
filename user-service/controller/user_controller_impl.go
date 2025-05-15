package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"user-service/exceptions"
	"user-service/helper"
	"user-service/model"
	"user-service/repository"
	"user-service/web/request"
	"user-service/web/response"
)

type UserControllerImpl struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
	OtpRepository  repository.OtpRepository
}

func NewUserController(
	DB *gorm.DB,
	userRepository repository.UserRepository,
	otpRepository repository.OtpRepository,
) *UserControllerImpl {
	return &UserControllerImpl{
		DB:             DB,
		UserRepository: userRepository,
		OtpRepository:  otpRepository,
	}
}

func (u UserControllerImpl) Profile(ctx *fiber.Ctx) error {
	claims := ctx.Locals("claims").(*helper.JWT)
	user, err := u.UserRepository.FindBy(u.DB, map[string]interface{}{
		"id": claims.Id,
	})
	if err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusNotFound, "user not found")
	}
	return ctx.Status(200).JSON(response.DataResponse("Get Profile Success", user, nil))
}

func (u UserControllerImpl) ProfileUpdate(ctx *fiber.Ctx) error {
	var data request.ProfileUpdateRequest
	claims := ctx.Locals("claims").(*helper.JWT)

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	dob, err := helper.ParseDate(data.DateOfBirth)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Tanggal lahir wajib format Y/m/d")
	}

	user, err := u.UserRepository.FindBy(u.DB, map[string]interface{}{
		"id": claims.Id,
	})
	if err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusNotFound, "user not found")
	}

	user.Name = data.Name
	user.DateOfBirth = dob
	if err = u.UserRepository.Update(u.DB, user); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Fail to update user")
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}

func (u UserControllerImpl) CheckToken(ctx *fiber.Ctx) error {
	claims := ctx.Locals("claims").(*helper.JWT)
	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OK Check Token", claims, nil))
}

func (u UserControllerImpl) LoginOrSignUp(ctx *fiber.Ctx) error {
	var data request.LoginOrSignupRequest

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	otpId, err := uuid.Parse(data.OtpId)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "otpId wajib format uuid")
	}

	phoneNumber, boolean := helper.ValidatePhoneNumber(data.PhoneNumber)

	if boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "No telepon wajib format +62 | 62 | 08")
	}

	var otp model.OTP
	otp.Id = otpId

	otp, boolean = u.OtpRepository.Verified(u.DB, otp)
	if boolean == false {
		return exceptions.ErrorHandlerBadRequest(ctx, "OTP Gagal atau tidak sesuai")
	}

	if otp.PhoneNumber != phoneNumber {
		return exceptions.ErrorHandlerBadRequest(ctx, "No telepon tidak sesuai dengan no telepon verifikasi OTP")
	}

	user := model.User{
		Id:          uuid.New(),
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}

	if err = u.UserRepository.Create(u.DB, user); err != nil {
		user, err = u.UserRepository.FindByPhoneNumber(u.DB, phoneNumber)

		if err != nil {
			return exceptions.ErrorHandlerBadRequest(ctx, "user tidak ditemukan")
		}
	}

	user.LastLoginAt = time.Now()
	if err = u.UserRepository.Update(u.DB, user); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Gagal update login terakhir")
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OK", map[string]interface{}{
		"id":    user.Id,
		"token": token,
		"name":  user.Name,
		"phone": user.PhoneNumber,
	}, nil))
}
