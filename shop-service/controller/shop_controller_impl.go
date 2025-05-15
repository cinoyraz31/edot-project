package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shop-service/exceptions"
	"shop-service/repository"
)

type ShopControllerImpl struct {
	DB             *gorm.DB
	ShopRepository repository.ShopRepository
}

func NewShopController(
	DB *gorm.DB,
	shopRepository repository.ShopRepository,
) *ShopControllerImpl {
	return &ShopControllerImpl{
		DB:             DB,
		ShopRepository: shopRepository,
	}
}

func (s ShopControllerImpl) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	shop, err := s.ShopRepository.FindBy(s.DB, map[string]interface{}{"id": id})

	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Shop not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(shop)
}
