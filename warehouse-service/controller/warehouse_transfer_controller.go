package controller

import "github.com/gofiber/fiber/v2"

type WarehouseTransferController interface {
	Add(ctx *fiber.Ctx) error
}
