package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"warehouse-service/controller"
	"warehouse-service/middleware"
	"warehouse-service/repository"
)

func WarehouseRoutes(app *fiber.App, db *gorm.DB) {
	warehouseRepository := repository.NewWarehouseRepository()
	warehouseController := controller.NewWarehouseController(db, warehouseRepository)

	app.Post("/warehouse", middleware.CheckTokenForShop, warehouseController.Add)
	app.Patch("/warehouse/:id", middleware.CheckTokenForShop, warehouseController.Edit)

	//	user
	app.Get("/user/shop/:shopId/product/:productId/warehouse", middleware.CheckToken, warehouseController.Show)
}
