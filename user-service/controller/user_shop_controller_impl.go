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

type UserShopControllerImpl struct {
	DB                 *gorm.DB
	OtpRepository      repository.OtpRepository
	UserShopRepository repository.UserShopRepository
}

func NewUserShopController(
	DB *gorm.DB,
	otpRepository repository.OtpRepository,
	userShopRepository repository.UserShopRepository,
) *UserShopControllerImpl {
	return &UserShopControllerImpl{
		DB:                 DB,
		OtpRepository:      otpRepository,
		UserShopRepository: userShopRepository,
	}
}

func (u UserShopControllerImpl) LoginOrSignUp(ctx *fiber.Ctx) error {
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

	userShop := model.UserShop{
		Id:          uuid.New(),
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}

	if err = u.UserShopRepository.Create(u.DB, userShop); err != nil {
		userShop, err = u.UserShopRepository.FindByPhoneNumber(u.DB, phoneNumber)

		if err != nil {
			return exceptions.ErrorHandlerBadRequest(ctx, "user tidak ditemukan")
		}
	}

	token, err := helper.GenerateTokenForShop(userShop)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OK", map[string]interface{}{
		"id":    userShop.Id,
		"token": token,
		"phone": userShop.PhoneNumber,
	}, nil))
}

func (u UserShopControllerImpl) CheckToken(ctx *fiber.Ctx) error {
	claims := ctx.Locals("claims").(*helper.JWTForShop)
	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("OK Check Token", claims, nil))
}
