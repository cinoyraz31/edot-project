package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"user-service/controller"
	"user-service/repository"
)

func OtpRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository()
	otpRepository := repository.NewOTPRepository()
	otpController := controller.NewOtpController(db, otpRepository, userRepository)

	app.Post("/time-count-otp", otpController.CountDown)
	app.Post("/send-otp", otpController.Send)
	app.Post("/validate-otp", otpController.Validate)
}
