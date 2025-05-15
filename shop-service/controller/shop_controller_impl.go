package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"shop-service/exceptions"
	warehouse_service "shop-service/external/warehouse-service"
	"shop-service/repository"
)

type ShopControllerImpl struct {
	DB                    *gorm.DB
	ShopRepository        repository.ShopRepository
	ShopProductRepository repository.ShopProductRepository
}

func NewShopController(
	DB *gorm.DB,
	shopRepository repository.ShopRepository,
	shopProductRepository repository.ShopProductRepository,
) *ShopControllerImpl {
	return &ShopControllerImpl{
		DB:                    DB,
		ShopRepository:        shopRepository,
		ShopProductRepository: shopProductRepository,
	}
}

func (s ShopControllerImpl) StockOnShops(ctx *fiber.Ctx) error {
	productId := ctx.Params("productId")

	shopProducts, err := s.ShopProductRepository.FindAll(s.DB, map[string]interface{}{
		"product_id":   productId,
		"shops.status": true,
	}, nil)
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "There issue get show products")
	}
	result := make([]interface{}, len(shopProducts))
	for i, shopProduct := range shopProducts {
		stockQtyResponse, _ := warehouse_service.StockQty(ctx, shopProduct.ShopID, shopProduct.ProductId)
		result[i] = map[string]interface{}{
			"shopId":    shopProduct.ShopID,
			"shopName":  shopProduct.ShopName,
			"productId": shopProduct.ProductId,
			"quantity":  stockQtyResponse.Quantity,
			"price":     shopProduct.Price,
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (s ShopControllerImpl) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	shop, err := s.ShopRepository.FindBy(s.DB, map[string]interface{}{"id": id})

	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Shop not found")
	}
	return ctx.Status(fiber.StatusOK).JSON(shop)
}
