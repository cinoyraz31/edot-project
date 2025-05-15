package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"user-service/controller"
	"user-service/middleware"
	"user-service/repository"
)

func UserShopRoutes(app *fiber.App, db *gorm.DB) {
	userShopRepository := repository.NewUserShopRepository()
	otpRepository := repository.NewOTPRepository()
	userShopController := controller.NewUserShopController(db, otpRepository, userShopRepository)

	app.Post("/shop/login-or-signup", userShopController.LoginOrSignUp)
	app.Get("/shop/check-token", middleware.JWTAuthenticateForShop, userShopController.CheckToken)
}
