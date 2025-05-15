package controller

import "github.com/gofiber/fiber/v2"

type OrderController interface {
	Add(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	Pay(ctx *fiber.Ctx) error
}
