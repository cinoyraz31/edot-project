package controller

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	Index(ctx *fiber.Ctx) error
}
