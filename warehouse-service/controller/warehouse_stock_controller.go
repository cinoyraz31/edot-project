package controller

import "github.com/gofiber/fiber/v2"

type WarehouseStockController interface {
	Add(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	StockOrder(ctx *fiber.Ctx) error
	StockRelease(ctx *fiber.Ctx) error
	StockPay(ctx *fiber.Ctx) error
	ProductQty(ctx *fiber.Ctx) error
}
