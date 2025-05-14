package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"shop-service/config"
	product_service "shop-service/external/product-service"
	"shop-service/model"
	"shop-service/repository"
	"sync"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

type Shop struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status bool   `json:"status"`
}

type ShopProduct struct {
	ShopCode    string  `json:"shopCode"`
	ProductCode string  `json:"productCode"`
	Price       float64 `json:"price"`
}

func JobForShopProduct(db *gorm.DB) {
	fileBytes, _ := ioutil.ReadFile("migration/shop_product.json")
	shopProductRepository := repository.NewShopProductRepository()
	shopRepository := repository.NewShopRepository()

	var shopProducts []ShopProduct
	if err := json.Unmarshal(fileBytes, &shopProducts); err != nil {
		log.Fatal(err)
	}
	for _, data := range shopProducts {
		//make sure check product to service product
		var wg sync.WaitGroup
		chanProduct := make(chan product_service.ExternalCheckProductResponse, 1)
		chanProductError := make(chan error, 1)
		chanShop := make(chan model.Shop, 1)
		chanShopError := make(chan error, 1)

		wg.Add(2)

		go func() {
			defer wg.Done()
			response, err := product_service.CheckExistProduct(data.ProductCode)
			chanProduct <- response
			chanProductError <- err
		}()

		go func() {
			defer wg.Done()
			shop, err := shopRepository.FindBy(db, map[string]interface{}{
				"code": data.ShopCode,
			})
			chanShop <- shop
			chanShopError <- err
		}()

		wg.Wait()
		close(chanProduct)
		close(chanProductError)
		close(chanShop)
		close(chanShopError)

		product := <-chanProduct
		productError := <-chanProductError
		shop := <-chanShop
		shopError := <-chanShopError

		if productError != nil {
			log.Fatal(productError)
			continue
		}
		if shopError != nil {
			log.Fatal(shopError)
			continue
		}

		shopProduct, err := shopProductRepository.FindBy(db, map[string]interface{}{
			"shop_id":    shop.Id,
			"product_id": product.Data.Id,
		})

		if err != nil {
			err := shopProductRepository.Create(db, model.ShopProduct{
				Id:        uuid.New(),
				ShopId:    shop.Id,
				ProductId: product.Data.Id,
				Price:     data.Price,
			})
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			shopProduct.Price = data.Price
			err := shopProductRepository.Update(db, shopProduct)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

	}
}

func JobForShop(db *gorm.DB) {
	fileBytes, _ := ioutil.ReadFile("migration/shop.json")
	repository := repository.NewShopRepository()

	var shops []Shop
	if err := json.Unmarshal(fileBytes, &shops); err != nil {
		log.Fatal(err)
	}
	for _, shop := range shops {
		_, err := repository.FindBy(db, map[string]interface{}{
			"code": shop.Code,
		})
		if err != nil {
			err := repository.Create(db, model.Shop{
				Id:     uuid.New(),
				Name:   shop.Name,
				Code:   shop.Code,
				Status: shop.Status,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	Init()
	db := config.OpenConnection()

	JobForShop(db)
	JobForShopProduct(db)
}
