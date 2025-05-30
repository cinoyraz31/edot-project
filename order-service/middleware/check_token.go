package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"order-service/exceptions"
	"os"
)

type CheckTokenResponse struct {
	Data struct {
		Id          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		PhoneNumber string    `json:"phoneNumber"`
	}
}

func CheckToken(ctx *fiber.Ctx) error {
	var response CheckTokenResponse
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/users/check-token", os.Getenv("API_GETAWAY_URL")), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Invalid token")
	}

	_ = json.Unmarshal(body, &response)
	ctx.Locals("claims", response)
	return ctx.Next()
}
