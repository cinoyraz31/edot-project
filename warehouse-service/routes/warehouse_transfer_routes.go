package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"warehouse-service/controller"
	"warehouse-service/middleware"
	"warehouse-service/repository"
)

func WarehouseTransferRoutes(app *fiber.App, db *gorm.DB) {
	transferRepository := repository.NewWarehouseTransferRepository()
	stockRepository := repository.NewWarehouseStockRepository()
	warehouseTransferController := controller.NewWarehouseTransferController(db, stockRepository, transferRepository)

	app.Post("/warehouse/transfer", middleware.CheckTokenForShop, warehouseTransferController.Add)
}
