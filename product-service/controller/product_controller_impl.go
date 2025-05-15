package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"product-service/exceptions"
	"product-service/repository"
	"product-service/web/request"
	"product-service/web/response"
)

type ProductControllerImpl struct {
	DB                *gorm.DB
	ProductRepository repository.ProductRepository
}

func NewProductController(
	DB *gorm.DB,
	productRepository repository.ProductRepository,
) *ProductControllerImpl {
	return &ProductControllerImpl{
		DB:                DB,
		ProductRepository: productRepository,
	}
}

func (p ProductControllerImpl) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := p.ProductRepository.FindBy(p.DB, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		product, err = p.ProductRepository.FindBy(p.DB, map[string]interface{}{
			"code": id,
		})
		if err != nil {
			return exceptions.ErrorHandlerBadRequest(ctx, err.Error())
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("product show", product, nil))
}

func (p ProductControllerImpl) Index(ctx *fiber.Ctx) error {
	data := request.Filter{
		Page: ctx.Query("page", "1"),
		Size: ctx.Query("size", "10"),
	}

	products, err := p.ProductRepository.FindAll(p.DB, map[string]interface{}{}, map[string]interface{}{
		"limit": data.Page,
		"size":  data.Size,
	})
	if err != nil {
		return exceptions.ErrorHandlerBadRequest(ctx, "Fail get products")
	}
	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("Get products", products, nil))
}
