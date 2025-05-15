package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"order-service/config"
	"order-service/exceptions"
	"order-service/routes"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	Init()
	db := config.OpenConnection()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(exceptions.ErrorHandlerInternalServerError)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to user service!")
	})

	routes.OrderRoutes(app, db)

	app.Use(func(ctx *fiber.Ctx) error {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusNotFound, "API Not Found")
	})

	err := app.Listen(os.Getenv("APP_URL"))
	if err != nil {
		panic(err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("Failed to close database connection!", err)
		}
		sqlDB.Close()
	}()
}
