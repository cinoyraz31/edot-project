package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"user-service/controller"
	"user-service/middleware"
	"user-service/repository"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repository.NewUserRepository()
	otpRepository := repository.NewOTPRepository()
	userController := controller.NewUserController(db, userRepository, otpRepository)

	app.Post("/login-or-signup", userController.LoginOrSignUp)
	app.Get("/check-token", middleware.JWTAuthenticate, userController.CheckToken)
}
