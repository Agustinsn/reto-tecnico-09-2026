package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/clients"
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/handlers"
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/services"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

const (
	defaultPort          = "8080"
	defaultNodeAPIURL    = "http://localhost:3000"
	defaultHTTPTimeoutMS = 5000
)

func main() {
	app := newApp()

	port := getEnv("PORT", defaultPort)

	log.Printf("Go API listening on :%s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}

func newApp() *fiber.App {
	qrService := services.NewQRService()

	nodeAPIURL := getEnv("NODE_API_URL", defaultNodeAPIURL)
	httpTimeout := getHTTPTimeout()

	nodeClient := clients.NewNodeClient(
		nodeAPIURL,
		httpTimeout,
	)

	qrWorkflow := services.NewQRWorkflow(
		qrService,
		nodeClient,
	)

	qrHandler := handlers.NewQRHandler(qrWorkflow)

	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New())

	api := app.Group("/api/v1")

	api.Get("/health", qrHandler.Health)
	api.Post("/qr", qrHandler.Calculate)

	return app
}

// getEnv returns the environment variable value or a default value.
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

// getHTTPTimeout returns the configured HTTP timeout for the Node.js API.
func getHTTPTimeout() time.Duration {
	value := os.Getenv("HTTP_TIMEOUT_MS")

	if value == "" {
		return defaultHTTPTimeoutMS * time.Millisecond
	}

	timeoutMS, err := strconv.Atoi(value)
	if err != nil || timeoutMS <= 0 {
		log.Printf(
			"Invalid HTTP_TIMEOUT_MS=%q, using default %dms",
			value,
			defaultHTTPTimeoutMS,
		)

		return defaultHTTPTimeoutMS * time.Millisecond
	}

	return time.Duration(timeoutMS) * time.Millisecond
}
