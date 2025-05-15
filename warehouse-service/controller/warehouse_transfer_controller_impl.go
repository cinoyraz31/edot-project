package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"warehouse-service/exceptions"
	"warehouse-service/repository"
	"warehouse-service/web/request/transfer"
)

type WarehouseTransferControllerImpl struct {
	DB                          *gorm.DB
	WarehouseStockRepository    repository.WarehouseStockRepository
	WarehouseTransferRepository repository.WarehouseTransferRepository
}

func NewWarehouseTransferController(
	DB *gorm.DB,
	warehouseStockRepository repository.WarehouseStockRepository,
	warehouseTransferRepository repository.WarehouseTransferRepository,
) *WarehouseTransferControllerImpl {
	return &WarehouseTransferControllerImpl{
		DB:                          DB,
		WarehouseStockRepository:    warehouseStockRepository,
		WarehouseTransferRepository: warehouseTransferRepository,
	}
}

func (w WarehouseTransferControllerImpl) Add(ctx *fiber.Ctx) error {
	var data transfer.CreateTransferRequest
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	fromStock, err := w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"warehouse_id": data.FromWarehouseId,
		"product_id":   data.ProductId,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock origin not found")
	}

	toStock, err := w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"warehouse_id": data.ToWarehouseId,
		"product_id":   data.ProductId,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock destination not found")
	}

	if (fromStock.Qty - fromStock.LockedQty) <= 0 {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock is empty")
	}

	fromStock.Qty -= data.Qty
	toStock.Qty += data.Qty

	_ = w.WarehouseStockRepository.Update(w.DB, fromStock)
	_ = w.WarehouseStockRepository.Update(w.DB, toStock)

	return ctx.Status(fiber.StatusCreated).JSON("")
}
