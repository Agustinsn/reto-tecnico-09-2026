package services

import (
	"fmt"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/clients"
	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
)

// QRWorkflow orchestrates QR factorization and communication with the Node.js API.
type QRWorkflow struct {
	qrService  *QRService
	nodeClient *clients.NodeClient
}

// NewQRWorkflow creates a new QR workflow.
func NewQRWorkflow(
	qrService *QRService,
	nodeClient *clients.NodeClient,
) *QRWorkflow {
	return &QRWorkflow{
		qrService:  qrService,
		nodeClient: nodeClient,
	}
}

// Execute calculates the QR factorization and sends the resulting
// matrices to the Node.js API for additional statistics.
func (w *QRWorkflow) Execute(matrix domain.Matrix) (*WorkflowResponse, error) {
	if w == nil {
		return nil, fmt.Errorf("QR workflow is nil")
	}

	if w.qrService == nil {
		return nil, fmt.Errorf("QR service is nil")
	}

	if w.nodeClient == nil {
		return nil, fmt.Errorf("Node client is nil")
	}

	q, r, err := w.qrService.Factorize(matrix)
	if err != nil {
		return nil, fmt.Errorf("QR factorization failed: %w", err)
	}

	statistics, err := w.nodeClient.CalculateStatistics(q, r)
	if err != nil {
		return nil, fmt.Errorf("Node API request failed: %w", err)
	}

	return &WorkflowResponse{
		Q:          q,
		R:          r,
		Statistics: *statistics,
	}, nil
}

// WorkflowResponse contains the complete result of the QR workflow.
type WorkflowResponse struct {
	Q          domain.Matrix              `json:"q"`
	R          domain.Matrix              `json:"r"`
	Statistics clients.StatisticsResponse `json:"statistics"`
}
