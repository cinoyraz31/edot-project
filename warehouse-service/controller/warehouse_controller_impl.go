package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"warehouse-service/exceptions"
	shop_service "warehouse-service/external/shop-service"
	"warehouse-service/model"
	"warehouse-service/repository"
	"warehouse-service/web/request/warehouse"
)

type WarehouseControllerImpl struct {
	DB                  *gorm.DB
	WarehouseRepository repository.WarehouseRepository
}

func NewWarehouseController(
	DB *gorm.DB,
	warehouseRepository repository.WarehouseRepository,
) *WarehouseControllerImpl {
	return &WarehouseControllerImpl{
		DB:                  DB,
		WarehouseRepository: warehouseRepository,
	}
}

func (w WarehouseControllerImpl) Add(ctx *fiber.Ctx) error {
	var data warehouse.CreateWarehouseRequest
	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}
	shopId, err := uuid.Parse(data.ShopId)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "shopId wajib format uuid")
	}
	shop, err := shop_service.ExternalShopDetail(ctx, shopId)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}

	if er := w.WarehouseRepository.Create(w.DB, model.Warehouse{
		Id:      uuid.New(),
		ShopId:  shop.Data.Id,
		Name:    data.Name,
		Address: data.Address,
		Status:  true,
	}); er != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}

func (w WarehouseControllerImpl) Edit(ctx *fiber.Ctx) error {
	var data warehouse.CreateWarehouseRequest
	id := ctx.Params("id")

	if err := ctx.BodyParser(&data); err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return exceptions.ErrorHandlerUnprocessableEntity(ctx, err)
	}

	shopId, err := uuid.Parse(data.ShopId)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "shopId wajib format uuid")
	}

	wareHouse, err := w.WarehouseRepository.FindBy(w.DB, map[string]interface{}{"id": id})
	if err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusNotFound, "warehouse not found")
	}
	_, err = shop_service.ExternalShopDetail(ctx, shopId)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
	}

	wareHouse.ShopId = shopId
	wareHouse.Name = data.Name
	wareHouse.Address = data.Address
	if err = w.WarehouseRepository.Update(w.DB, wareHouse); err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusNoContent).JSON("")
}
