package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"shop-service/config"
	"shop-service/exceptions"
	"shop-service/routes"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	if err := config.Sentry(); err != nil {
		panic(fmt.Sprintf("sentry.Init: %s", err))
	}
}

func main() {
	Init()
	db := config.OpenConnection()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(exceptions.ErrorHandlerInternalServerError)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to shop service!")
	})

	routes.ShopRoutes(app, db)

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
