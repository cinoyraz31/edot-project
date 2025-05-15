package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"product-service/controller"
	"product-service/middleware"
	"product-service/repository"
)

func ProductRoutes(app *fiber.App, db *gorm.DB) {
	productRepository := repository.NewProductRepository()
	productController := controller.NewProductController(db, productRepository)

	app.Get("products", middleware.CheckToken, productController.Index)

	//	internal
	app.Get("static/products/:id", middleware.StaticToken, productController.Show)
}
