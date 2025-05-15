package controller

import "github.com/gofiber/fiber/v2"

type WarehouseController interface {
	Add(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
}
