package controller

import "github.com/gofiber/fiber/v2"

type UserShopController interface {
	LoginOrSignUp(ctx *fiber.Ctx) error
	CheckToken(ctx *fiber.Ctx) error
}
