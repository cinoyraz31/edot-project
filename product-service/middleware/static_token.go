package middleware

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"product-service/exceptions"
	"strings"
)

func StaticToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	if tokenString != os.Getenv("INTERNAL_TOKEN") {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Invalid token")
	}
	return ctx.Next()
}
