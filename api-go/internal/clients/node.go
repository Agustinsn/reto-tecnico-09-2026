package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
)

const defaultHTTPTimeout = 5 * time.Second

// NodeClient handles HTTP communication with the Node.js API.
type NodeClient struct {
	baseURL    string
	httpClient *http.Client
}

// StatisticsRequest is the payload sent from the Go API to the Node.js API.
type StatisticsRequest struct {
	Q domain.Matrix `json:"q"`
	R domain.Matrix `json:"r"`
}

// MatrixStatistics contains the statistics calculated by the Node.js API.
type MatrixStatistics struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Average    float64 `json:"average"`
	Sum        float64 `json:"sum"`
	IsDiagonal bool    `json:"isDiagonal"`
}

// StatisticsResponse is the response returned by the Node.js API.
type StatisticsResponse struct {
	Q           MatrixStatistics `json:"q"`
	R           MatrixStatistics `json:"r"`
	AnyDiagonal bool             `json:"anyDiagonal"`
}

// NewNodeClient creates a client for communicating with the Node.js API.
func NewNodeClient(baseURL string, timeout time.Duration) *NodeClient {
	baseURL = strings.TrimRight(baseURL, "/")

	if timeout <= 0 {
		timeout = defaultHTTPTimeout
	}

	return &NodeClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// CalculateStatistics sends the Q and R matrices to the Node.js API
// and returns the calculated statistics.
func (c *NodeClient) CalculateStatistics(
	q domain.Matrix,
	r domain.Matrix,
) (*StatisticsResponse, error) {
	if c == nil {
		return nil, errors.New("node client is nil")
	}

	if c.baseURL == "" {
		return nil, errors.New("node API URL is empty")
	}

	requestBody := StatisticsRequest{
		Q: q,
		R: r,
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Node API request: %w", err)
	}

	url := c.baseURL + "/api/v1/statistics"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create Node API request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call Node API: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Node API response: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf(
			"Node API returned status %d: %s",
			resp.StatusCode,
			strings.TrimSpace(string(responseBody)),
		)
	}

	var statistics StatisticsResponse

	if err := json.Unmarshal(responseBody, &statistics); err != nil {
		return nil, fmt.Errorf("failed to decode Node API response: %w", err)
	}

	return &statistics, nil
}
