package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	LoginOrSignUp(ctx *fiber.Ctx) error
	CheckToken(ctx *fiber.Ctx) error
	Profile(ctx *fiber.Ctx) error
	ProfileUpdate(ctx *fiber.Ctx) error
}
