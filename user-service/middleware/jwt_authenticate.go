package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
	"user-service/exceptions"
	"user-service/helper"
)

func JWTAuthenticateForShop(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &helper.JWTForShop{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_FOR_SHOP")), nil
	})
	if err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(*helper.JWTForShop)
	if !ok || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	ctx.Locals("claims", claims)
	return ctx.Next()
}

func JWTAuthenticate(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenString, &helper.JWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return exceptions.ErrorHandlerCustom(ctx, fiber.StatusUnauthorized, "Invalid token")
	}

	claims, ok := token.Claims.(*helper.JWT)
	if !ok || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	ctx.Locals("claims", claims)
	return ctx.Next()
}
