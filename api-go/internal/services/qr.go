package services

import (
	"errors"
	"math"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
)

var (
	ErrEmptyMatrix    = errors.New("matrix must not be empty")
	ErrEmptyRow       = errors.New("matrix rows must not be empty")
	ErrNonRectangular = errors.New("matrix must be rectangular")
)

// QRService calculates the reduced QR factorization A = Q * R using
// Householder reflections. For an m x n matrix, Q is m x k and R is k x n,
// where k = min(m, n).
type QRService struct{}

func NewQRService() *QRService {
	return &QRService{}
}

func (s *QRService) Factorize(a domain.Matrix) (domain.Matrix, domain.Matrix, error) {
	rows, cols, err := validateMatrix(a)
	if err != nil {
		return nil, nil, err
	}

	k := min(rows, cols)
	r := cloneMatrix(a)
	qFull := identityMatrix(rows)

	for j := 0; j < k; j++ {
		x := make([]float64, rows-j)
		for i := j; i < rows; i++ {
			x[i-j] = r[i][j]
		}

		norm := vectorNorm(x)
		if norm < 1e-15 {
			continue
		}

		alpha := -math.Copysign(norm, x[0])
		v := append([]float64(nil), x...)
		v[0] -= alpha

		vNormSq := dot(v, v)
		if vNormSq < 1e-30 {
			continue
		}

		beta := 2.0 / vNormSq

		// Apply H = I - beta*v*v^T to R, from the left.
		for col := j; col < cols; col++ {
			projection := 0.0
			for i := j; i < rows; i++ {
				projection += v[i-j] * r[i][col]
			}
			projection *= beta

			for i := j; i < rows; i++ {
				r[i][col] -= projection * v[i-j]
			}
		}

		// Q is the product H1*H2*...*Hk. Since H is symmetric,
		// multiplying the current Q by H on the right means updating columns.
		for row := 0; row < rows; row++ {
			projection := 0.0
			for col := j; col < rows; col++ {
				projection += qFull[row][col] * v[col-j]
			}
			projection *= beta

			for col := j; col < rows; col++ {
				qFull[row][col] -= projection * v[col-j]
			}
		}
	}

	q := make(domain.Matrix, rows)
	for i := range q {
		q[i] = make([]float64, k)
		copy(q[i], qFull[i][:k])
	}

	// Clean very small floating-point noise below the diagonal of R.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > j && math.Abs(r[i][j]) < 1e-12 {
				r[i][j] = 0
			}
		}
	}

	rReduced := make(domain.Matrix, k)
	for i := 0; i < k; i++ {
		rReduced[i] = make([]float64, cols)
		copy(rReduced[i], r[i])
	}

	return q, rReduced, nil
}

func validateMatrix(a domain.Matrix) (int, int, error) {
	if len(a) == 0 {
		return 0, 0, ErrEmptyMatrix
	}
	if len(a[0]) == 0 {
		return 0, 0, ErrEmptyRow
	}

	cols := len(a[0])
	for i := 1; i < len(a); i++ {
		if len(a[i]) == 0 {
			return 0, 0, ErrEmptyRow
		}
		if len(a[i]) != cols {
			return 0, 0, ErrNonRectangular
		}
	}

	return len(a), cols, nil
}

func cloneMatrix(a domain.Matrix) domain.Matrix {
	out := make(domain.Matrix, len(a))
	for i := range a {
		out[i] = append([]float64(nil), a[i]...)
	}
	return out
}

func identityMatrix(n int) domain.Matrix {
	m := make(domain.Matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		m[i][i] = 1
	}
	return m
}

func vectorNorm(v []float64) float64 {
	sum := 0.0
	for _, x := range v {
		sum += x * x
	}
	return math.Sqrt(sum)
}

func dot(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
