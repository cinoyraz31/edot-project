package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"product-service/exceptions"
	shop_service "product-service/external/shop-service"
	"product-service/helper"
	"product-service/model"
	"product-service/repository"
	"product-service/web/request"
	"product-service/web/response"
	"strconv"
	"sync"
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

	var wg sync.WaitGroup
	chanCount := make(chan int64, 1)
	chanProducts := make(chan []model.Product, 1)

	wg.Add(2)

	go func() {
		defer wg.Done()
		intPage, _ := strconv.Atoi(fmt.Sprintf("%v", data.Page))
		intPerPage, _ := strconv.Atoi(fmt.Sprintf("%v", data.Size))

		products, _ := p.ProductRepository.FindAll(p.DB, map[string]interface{}{
			"status": true,
		}, map[string]interface{}{
			"offset": (intPage - 1) * intPerPage,
			"limit":  intPerPage,
		})
		chanProducts <- products
	}()

	go func() {
		defer wg.Done()
		count, _ := p.ProductRepository.Count(p.DB, map[string]interface{}{
			"status": true,
		})
		chanCount <- count
	}()

	wg.Wait()
	close(chanCount)
	close(chanProducts)

	products := <-chanProducts
	count := <-chanCount
	pagination := helper.MakePagination(count, data.Page, data.Size)

	result := make([]interface{}, len(products))
	for i, product := range products {
		shopsResponse, _ := shop_service.ExternalProductShopList(ctx, product.Id)
		result[i] = map[string]interface{}{
			"id":          product.Id,
			"name":        product.Name,
			"code":        product.Code,
			"description": product.Description,
			"shops":       shopsResponse,
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponse("Get products", result, map[string]interface{}{
		"pagination": pagination,
	}))
}
