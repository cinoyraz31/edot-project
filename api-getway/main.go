package main

import (
	"api-getway/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	Init()
	app := fiber.New()

	app.All("/api/users/*", controller.Proxy(os.Getenv("USER_SERVICE"), "/api/users"))
	app.All("/api/products/*", controller.Proxy(os.Getenv("PRODUCT_SERVICE"), "/api/products"))
	app.All("/api/shops/*", controller.Proxy(os.Getenv("SHOP_SERVICE"), "/api/shops"))
	app.All("/api/warehouses/*", controller.Proxy(os.Getenv("WAREHOUSE_SERVICE"), "/api/warehouses"))
	app.All("/api/orders/*", controller.Proxy(os.Getenv("ORDER_SERVICE"), "/api/orders"))

	err := app.Listen(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}
}
