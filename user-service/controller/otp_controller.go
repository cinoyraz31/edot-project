package controller

import "github.com/gofiber/fiber/v2"

type OtpController interface {
	Send(ctx *fiber.Ctx) error
	Validate(ctx *fiber.Ctx) error
	CountDown(ctx *fiber.Ctx) error
}
