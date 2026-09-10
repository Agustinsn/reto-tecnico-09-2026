package domain

// Matrix represents a rectangular matrix of float64 values.
type Matrix [][]float64

// QRRequest is the input contract for the QR endpoint.
type QRRequest struct {
	Matrix Matrix `json:"matrix"`
}

// QRResponse contains the reduced QR factorization of the input matrix.
type QRResponse struct {
	Q Matrix `json:"q"`
	R Matrix `json:"r"`
}
