package handlers

import (
	"github.com/gofiber/fiber/v3"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var request LoginRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.authService.Login(request.Username, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
