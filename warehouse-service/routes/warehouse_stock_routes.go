package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"warehouse-service/controller"
	"warehouse-service/middleware"
	"warehouse-service/repository"
)

func WarehouseStockRoutes(app *fiber.App, db *gorm.DB) {
	stockRepository := repository.NewWarehouseStockRepository()
	warehouseRepository := repository.NewWarehouseRepository()
	warehouseStockController := controller.NewWarehouseStockController(db, stockRepository, warehouseRepository)

	app.Post("/warehouse/stock", middleware.CheckTokenForShop, warehouseStockController.Add)
	app.Patch("/warehouse/stock/:id", middleware.CheckTokenForShop, warehouseStockController.Edit)
	app.Delete("/warehouse/stock/:id", middleware.CheckTokenForShop, warehouseStockController.Delete)

	//	user
	app.Post("/user/warehouse-stock/order", middleware.CheckToken, warehouseStockController.StockOrder)
	app.Get("/user/shop/:shopId/product/:productId/warehouse-stock", middleware.CheckToken, warehouseStockController.ProductQty)
}
