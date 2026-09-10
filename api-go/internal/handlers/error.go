package handlers

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(c fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := "internal server error"

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		status = fiberErr.Code
		message = fiberErr.Message
	} else if err != nil {
		log.Printf("request failed: %v", err)
	}

	return c.Status(status).JSON(ErrorResponse{Error: message})
}
