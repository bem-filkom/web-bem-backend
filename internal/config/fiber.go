package config

import (
	"errors"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/response"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
)

func NewFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName:      "web-bem-backend",
			JSONEncoder:  jsoniter.Marshal,
			JSONDecoder:  jsoniter.Unmarshal,
			ErrorHandler: newErrorHandler(),
		})

	app.Use(logger.New())
	return app
}

func newErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		var apiErr *response.ErrorResponse
		if errors.As(err, &apiErr) {
			return ctx.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		var validationErr *validator.ValidationErrorsResponse
		if errors.As(err, &validationErr) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Validation error",
				"detail":  validationErr,
			})
		}

		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"message": fiberErr.Message,
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
}
