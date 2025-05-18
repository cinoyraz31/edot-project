package exceptions

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

type ErrorHandlerResp struct {
	Message string      `json:"message"`
	Err     interface{} `json:"errors"`
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorHandlerInternalServerError(ctx *fiber.Ctx) error {
	defer func() {
		sentryIsActive, _ := strconv.ParseBool(os.Getenv("SENTRY_ENABLED"))

		if r := recover(); r != nil {
			// capture sentry
			if sentryIsActive {
				sentry.CaptureException(errors.New(fmt.Sprintf("%v", r)))
			}

			ErrorHandlerCustom(ctx, fiber.StatusInternalServerError, fmt.Sprintf("%v", r))
		}
	}()
	return ctx.Next()
}

func ErrorHandlerUnprocessableEntity(ctx *fiber.Ctx, err error) error {
	validationErrors := err.(validator.ValidationErrors)
	errorArr := make([]ErrorField, len(validationErrors))

	for i, err := range validationErrors {
		errorArr[i] = ErrorField{
			Field:   err.Field(),
			Message: err.Tag(),
		}
	}

	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ErrorHandlerResp{
		Message: err.Error(),
		Err:     errorArr,
	})
}

func ErrorHandlerCustom(ctx *fiber.Ctx, code int, message string) error {
	return ctx.Status(code).JSON(ErrorHandlerResp{
		Message: message,
	})
}

func ErrorHandlerBadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(ErrorHandlerResp{
		Message: message,
	})
}
