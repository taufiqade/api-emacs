package main

import (
	"api-survey-go/config"
	"api-survey-go/internal/model"
	"api-survey-go/pkg/utils"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	fmt.Println("-----error handler", code, err)
	return c.Status(code).JSON(model.Response{
		Status:  "error",
		Message: err.Error(),
	})
}

// Middleware to parse and validate JWT token
var parseJWT = func(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return c.Next()
	}

	// Split the Authorization header value by space
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header format")
	}

	// Extract token string without the "Bearer" prefix
	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}

	c.Locals("auth", &model.Auth{
		ID:       utils.ConvertToString(claims["user_id"]),
		RoleID:   utils.ConvertToString(claims["role_id"]),
		ClientID: utils.ConvertToString(claims["client_id"]),
	})

	return c.Next()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})

	app.Use(parseJWT)

	log := config.NewLogger()
	validate := config.NewValidator()
	db := config.NewDB(&config.DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		RedisURL:   os.Getenv("REDIS_URL"),
	})

	config.Bootstrap(&config.AppConfig{
		App:      app,
		Log:      log,
		DB:       db,
		Validate: validate,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
