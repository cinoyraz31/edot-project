package shop_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
)

func ExternalProductShopList(ctx *fiber.Ctx, productId uuid.UUID) ([]interface{}, error) {
	var response []interface{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/shops/user/product/%s/shops", os.Getenv("API_GETAWAY_URL"), productId), nil)
	if err != nil {
		return response, err
	}
	req.Header.Set("Authorization", ctx.Get("Authorization"))
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
		return response, errors.New("Shop not found")
	}

	_ = json.Unmarshal(body, &response)
	return response, nil
}
