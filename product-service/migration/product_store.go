package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"product-service/config"
	"product-service/model"
	"product-service/repository"
)

type ProductStore struct {
	Name        string
	Code        string
	Description string
	Status      bool
}

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	Init()
	db := config.OpenConnection()
	fileBytes, _ := ioutil.ReadFile("migration/product_store.json")

	productRepository := repository.NewProductRepository()

	var products []ProductStore
	if err := json.Unmarshal(fileBytes, &products); err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		_, err := productRepository.FindBy(db, map[string]interface{}{
			"code": product.Code,
		})

		if err != nil {
			err := productRepository.Create(db, model.Product{
				Id:          uuid.New(),
				Name:        product.Name,
				Code:        product.Code,
				Description: product.Description,
				Status:      product.Status,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
