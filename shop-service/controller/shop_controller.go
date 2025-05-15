package controller

import "github.com/gofiber/fiber/v2"

type ShopController interface {
	Show(ctx *fiber.Ctx) error
}
