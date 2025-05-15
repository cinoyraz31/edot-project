package product_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

type ExternalProductResponse struct {
	Data struct {
		Id          string `json:"id"`
		PhoneNumber string `json:"phoneNumber"`
	}
}

func ExternalProductDetail(productId uuid.UUID) (ExternalProductResponse, error) {
	var response ExternalProductResponse
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/static/products/%s", os.Getenv("PRODUCT_URL"), productId), nil)
	if err != nil {
		return response, err
	}
	req.Header.Set("Authorization", os.Getenv("INTERNAL_TOKEN"))
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

	if resp.StatusCode != http.StatusOK {
		return response, errors.New("product not found")
	}

	_ = json.Unmarshal(body, &response)
	return response, nil
}
