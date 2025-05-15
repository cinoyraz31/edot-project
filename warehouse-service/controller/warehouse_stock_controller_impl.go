package controller

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"warehouse-service/exceptions"
	product_service "warehouse-service/external/product-service"
	"warehouse-service/model"
	"warehouse-service/repository"
	"warehouse-service/web/request/stock"
)

type WarehouseStockControllerImpl struct {
	DB                       *gorm.DB
	WarehouseStockRepository repository.WarehouseStockRepository
	WarehouseRepository      repository.WarehouseRepository
}

func NewWarehouseStockController(
	DB *gorm.DB,
	warehouseStockRepository repository.WarehouseStockRepository,
	warehouseRepository repository.WarehouseRepository,
) *WarehouseStockControllerImpl {
	return &WarehouseStockControllerImpl{
		DB:                       DB,
		WarehouseStockRepository: warehouseStockRepository,
		WarehouseRepository:      warehouseRepository,
	}
}

func checkExistWarehouseAndProduct(w WarehouseStockControllerImpl, data stock.CreateStockRequest) error {
	_, err := w.WarehouseRepository.FindBy(w.DB, map[string]interface{}{
		"id": data.WarehouseId,
	})
	if err != nil {
		return errors.New("Warehouse not found")
	}

	_, err = product_service.ExternalProductDetail(data.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (w WarehouseStockControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	stock, err := w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "stock not found")
	}

	if stock.LockedQty > 0 {
		return exceptions.ErrorHandlerBadRequest(ctx, "There is still stock that is being processed order")
	}

	if err := w.WarehouseStockRepository.Delete(w.DB, stock.Id); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "delete stock is problem")
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}

func (w WarehouseStockControllerImpl) Add(ctx *fiber.Ctx) error {
	var data stock.CreateStockRequest
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	err = checkExistWarehouseAndProduct(w, data)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}

	_, err = w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"warehouse_id": data.WarehouseId,
		"product_id":   data.ProductId,
	})
	if err == nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Warehouse already exists")
	}

	if err := w.WarehouseStockRepository.Create(w.DB, model.Stock{
		Id:          uuid.New(),
		WarehouseId: data.WarehouseId,
		ProductId:   data.ProductId,
		Qty:         data.Qty,
	}); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Failed create stock")
	}
	return ctx.Status(fiber.StatusCreated).JSON("")
}

func (w WarehouseStockControllerImpl) Edit(ctx *fiber.Ctx) error {
	var data stock.UpdateStockRequest
	id := ctx.Params("id")

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	stock, err := w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock not found")
	}

	if data.Qty < stock.LockedQty {
		return exceptions.ErrorHandlerBadRequest(ctx, "There is still stock that is being processed, wait for the order to be paid or cancelled")
	}
	stock.Qty = data.Qty
	if err := w.WarehouseStockRepository.Update(w.DB, stock); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Failed update stock")
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}

func (w WarehouseStockControllerImpl) StockOrder(ctx *fiber.Ctx) error {
	var data stock.OrderStockRequest
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}
	stock, err := w.WarehouseStockRepository.FindBy(w.DB, map[string]interface{}{
		"warehouse_id": data.WarehouseId,
		"product_id":   data.ProductId,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock not found")
	}

	if (stock.Qty - stock.LockedQty) <= 0 {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock is empty")
	}

	if data.Qty > (stock.Qty - stock.LockedQty) {
		return exceptions.ErrorHandlerBadRequest(ctx, "Stock not enough")
	}
	if err := w.WarehouseStockRepository.LockQty(w.DB, data.WarehouseId, data.ProductId, data.Qty); err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Is problem lock stock")
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}
