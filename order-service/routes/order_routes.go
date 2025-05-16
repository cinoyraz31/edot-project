package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"order-service/controller"
	"order-service/middleware"
	"order-service/repository"
)

func OrderRoutes(app *fiber.App, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository()
	shopOrderRepo := repository.NewShopOrderRepository()
	orderItemRepo := repository.NewOrderItemRepository()
	orderController := controller.NewOrderController(orderItemRepo, db, orderRepo, shopOrderRepo)

	app.Post("/orders", middleware.CheckToken, orderController.Add)
	app.Get("/orders", middleware.CheckToken, orderController.List)
	app.Get("/orders/:orderNumber", middleware.CheckToken, orderController.Show)
	app.Post("/webhook/order-pay", middleware.WebhookPayment, orderController.Pay)
}
