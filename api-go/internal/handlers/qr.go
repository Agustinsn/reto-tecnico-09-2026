package handlers

import (
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/services"
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/validators"
	"github.com/gofiber/fiber/v3"
)

// QRHandler handles HTTP requests related to QR factorization.
type QRHandler struct {
	workflow *services.QRWorkflow
}

// NewQRHandler creates a new QR handler.
func NewQRHandler(workflow *services.QRWorkflow) *QRHandler {
	return &QRHandler{
		workflow: workflow,
	}
}

// Calculate validates the request, calculates the QR factorization,
// sends the result to the Node.js API, and returns the complete response.
func (h *QRHandler) Calculate(c fiber.Ctx) error {
	var request domain.QRRequest

	if err := c.Bind().JSON(&request); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid JSON request",
		)
	}

	if err := validators.Validate(request.Matrix); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	if h == nil || h.workflow == nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"QR workflow is not configured",
		)
	}

	result, err := h.workflow.Execute(request.Matrix)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

// Health returns the health status of the Go API.
func (h *QRHandler) Health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}