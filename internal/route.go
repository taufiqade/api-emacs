package internal

import (
	"api-survey-go/internal/handler"
	"api-survey-go/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App         *fiber.App
	UserHandler *handler.UserHandler
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")

	// guest api routes
	guest := api.Group("/")
	guest.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("pong")
	})
	guest.Post("/login", c.UserHandler.Login)

	// authenticated route
	user := api.Group("/user", middleware.Protected())
	user.Get("/:id", middleware.OnlyAdmin(), c.UserHandler.FindById)
	user.Post("/", c.UserHandler.Create)
	user.Put("/", c.UserHandler.Update)
	user.Get("/", c.UserHandler.GetCurrentUser)
}
