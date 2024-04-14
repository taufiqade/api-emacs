package middleware

import (
	"api-survey-go/internal/model"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ErrorHandler: jwtError,
	})
}

func OnlyAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Locals("auth").(model.Auth)
		if auth.RoleID != "1" {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		}
		return c.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(model.Response{
				Status:  "error",
				Message: "missing or malformed JWT",
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(model.Response{Status: "error", Message: "Invalid or expired JWT"})
}
