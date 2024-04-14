package config

import (
	"api-survey-go/internal"
	"api-survey-go/internal/handler"
	"api-survey-go/internal/repository"
	"api-survey-go/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppConfig struct {
	Log      *logrus.Logger
	App      *fiber.App
	DB       *gorm.DB
	Validate *validator.Validate
}

func Bootstrap(config *AppConfig) {
	userRepo := repository.NewUserRepository(config.DB)

	userService := service.NewUserService(config.Log, userRepo, config.Validate, config.DB)

	userHandler := handler.NewUserHandler(userService, config.Log)

	routeConfig := internal.RouteConfig{
		App:         config.App,
		UserHandler: userHandler,
	}

	routeConfig.Setup()
}
