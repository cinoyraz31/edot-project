package middleware

import (
	"github.com/gofiber/fiber/v2"
	"order-service/exceptions"
	"os"
	"strings"
)

func WebhookPayment(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	if tokenString != os.Getenv("WEBHOOK_PAYMENT_TOKEN") {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Invalid token")
	}
	return ctx.Next()
}
