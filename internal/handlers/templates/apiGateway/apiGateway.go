package apiGateway

import (
	"github.com/Hell077/GoBoost/internal/utils"
)

func CreateApiGatewayTemplate(projectDir, projectName string) {
	directories := []string{
		"cmd/app",
		"internal/config",
		"internal/handlers",
		"internal/routes",
		"internal/service",
		"internal/middleware",
		"internal/utils",
		"pkg/logger",
	}

	files := map[string]string{
		"cmd/app/main.go": `package main

import (
	"log"
	"api-gateway/internal/config"
	"api-gateway/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Create a new Fiber application
	app := fiber.New()

	// Set up routing
	routes.SetupRoutes(app)

	// Start server on port 8080
	log.Fatal(app.Listen(":8080"))
}
`,
		"internal/config/config.go": `package config

import (
	"log"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
`,
		"internal/handlers/handler.go": `package handlers

import "github.com/gofiber/fiber/v2"

// Handler for Service1
func Service1Handler(c *fiber.Ctx) error {
	return c.SendString("Service 1 Handler")
}

// Handler for Service2
func Service2Handler(c *fiber.Ctx) error {
	return c.SendString("Service 2 Handler")
}
`,
		"internal/routes/router.go": `package routes

import (
	"api-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// Set up routing
func SetupRoutes(app *fiber.App) {
	app.Get("/service1", handlers.Service1Handler) // GET request for service 1
	app.Post("/service2", handlers.Service2Handler) // POST request for service 2
}
`,
		"internal/service/service.go": `package service

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func CallService1() (string, error) {
	// Example call to service 1
	resp, err := http.Get("http://service1/api")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("service1 error")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func CallService2() (string, error) {
	// Example call to service 2
	resp, err := http.Post("http://service2/api", "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("service2 error")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
`,
		"internal/middleware/auth.go": `package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Simple authentication example
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	// Token validation logic
	if !ValidateToken(token) {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	return c.Next()
}

func ValidateToken(token string) bool {
	// Simple token validation
	return token == "valid-token"
}
`,
		"internal/utils/utils.go": `package utils

func UtilityFunction() {
	// General utilities
}
`,
		"pkg/logger/logger.go": `package logger

import "log"

func InitLogger() {
	// Logger initialization
	log.Println("Logger initialized")
}
`,
		".gitignore":     utils.Ignore(),
		".gitattributes": utils.Attribute(),
	}

	utils.CreateProjectStructure(projectDir, directories, files, projectName)
}
