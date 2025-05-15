package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shop-service/controller"
	"shop-service/middleware"
	"shop-service/repository"
)

func ShopRoutes(app *fiber.App, db *gorm.DB) {
	shopRepository := repository.NewShopRepository()
	shopController := controller.NewShopController(db, shopRepository)

	app.Get("/internal/shops/:id", middleware.CheckTokenForShop, shopController.Show)
}
