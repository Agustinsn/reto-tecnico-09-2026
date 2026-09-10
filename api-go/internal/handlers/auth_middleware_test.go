package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/services"
)

func TestJWTMiddlewareWithoutAuthorizationHeader(t *testing.T) {
	authService := services.NewAuthService("admin", "admin123", "test-secret")

	app := fiber.New()
	app.Use(JWTMiddleware(authService))
	app.Get("/protected", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	request := httptest.NewRequest("GET", "/protected", nil)

	response, err := app.Test(request)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", response.StatusCode)
	}
}

func TestJWTMiddlewareWithInvalidToken(t *testing.T) {
	authService := services.NewAuthService("admin", "admin123", "test-secret")

	app := fiber.New()
	app.Use(JWTMiddleware(authService))
	app.Get("/protected", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	request := httptest.NewRequest("GET", "/protected", nil)
	request.Header.Set("Authorization", "Bearer invalid-token")

	response, err := app.Test(request)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.StatusCode != fiber.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", response.StatusCode)
	}
}

func TestJWTMiddlewareWithValidToken(t *testing.T) {
	authService := services.NewAuthService("admin", "admin123", "test-secret")

	token, err := authService.Login("admin", "admin123")
	if err != nil {
		t.Fatalf("expected no error generating token, got %v", err)
	}

	app := fiber.New()
	app.Use(JWTMiddleware(authService))
	app.Get("/protected", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	request := httptest.NewRequest("GET", "/protected", nil)
	request.Header.Set("Authorization", "Bearer "+token)

	response, err := app.Test(request)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status 200, got %d", response.StatusCode)
	}
}