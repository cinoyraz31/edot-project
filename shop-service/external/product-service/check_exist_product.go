package product_service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type ExternalCheckProductResponse struct {
	Data struct {
		Id uuid.UUID `json:"id"`
	} `json:"data"`
}

func CheckExistProduct(productCode string) (ExternalCheckProductResponse, error) {
	var response ExternalCheckProductResponse
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/products/static/products/%s", os.Getenv("API_GETAWAY_URL"), productCode), nil)
	if err != nil {
		return response, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("INTERNAL_TOKEN")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	_ = json.Unmarshal(body, &response)
	return response, nil
}
