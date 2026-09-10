package services

import (
	"math"
	"testing"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
)

func TestFactorizeTallMatrix(t *testing.T) {
	a := domain.Matrix{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	service := NewQRService()
	q, r, err := service.Factorize(a)
	if err != nil {
		t.Fatalf("Factorize() error = %v", err)
	}

	assertShape(t, q, 3, 2)
	assertShape(t, r, 2, 2)
	assertMatrixApprox(t, multiply(q, r), a, 1e-10)
	assertOrthogonalColumns(t, q, 1e-10)
}

func TestFactorizeSquareMatrix(t *testing.T) {
	a := domain.Matrix{
		{1, 2},
		{3, 4},
	}

	service := NewQRService()
	q, r, err := service.Factorize(a)
	if err != nil {
		t.Fatalf("Factorize() error = %v", err)
	}

	assertShape(t, q, 2, 2)
	assertShape(t, r, 2, 2)
	assertMatrixApprox(t, multiply(q, r), a, 1e-10)
}

func TestFactorizeWideMatrix(t *testing.T) {
	a := domain.Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	service := NewQRService()
	q, r, err := service.Factorize(a)
	if err != nil {
		t.Fatalf("Factorize() error = %v", err)
	}

	assertShape(t, q, 2, 2)
	assertShape(t, r, 2, 3)
	assertMatrixApprox(t, multiply(q, r), a, 1e-10)
}

func TestFactorizeRejectsNonRectangular(t *testing.T) {
	a := domain.Matrix{
		{1, 2, 3},
		{4, 5},
	}

	service := NewQRService()
	_, _, err := service.Factorize(a)
	if err != ErrNonRectangular {
		t.Fatalf("expected ErrNonRectangular, got %v", err)
	}
}

func multiply(a, b domain.Matrix) domain.Matrix {
	rows := len(a)
	inner := len(b)
	cols := len(b[0])
	out := make(domain.Matrix, rows)

	for i := range out {
		out[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			for k := 0; k < inner; k++ {
				out[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return out
}

func assertMatrixApprox(t *testing.T, got, want domain.Matrix, tolerance float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("row count mismatch: got %d want %d", len(got), len(want))
	}
	for i := range got {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("column count mismatch on row %d", i)
		}
		for j := range got[i] {
			if math.Abs(got[i][j]-want[i][j]) > tolerance {
				t.Fatalf("matrix mismatch at [%d][%d]: got %.15f want %.15f", i, j, got[i][j], want[i][j])
			}
		}
	}
}

func assertOrthogonalColumns(t *testing.T, q domain.Matrix, tolerance float64) {
	t.Helper()
	cols := len(q[0])
	for i := 0; i < cols; i++ {
		for j := 0; j < cols; j++ {
			dotProduct := 0.0
			for row := range q {
				dotProduct += q[row][i] * q[row][j]
			}
			expected := 0.0
			if i == j {
				expected = 1
			}
			if math.Abs(dotProduct-expected) > tolerance {
				t.Fatalf("Q is not orthonormal at [%d][%d]: got %.15f", i, j, dotProduct)
			}
		}
	}
}

func assertShape(t *testing.T, matrix domain.Matrix, rows, cols int) {
	t.Helper()
	if len(matrix) != rows {
		t.Fatalf("rows: got %d want %d", len(matrix), rows)
	}
	for i := range matrix {
		if len(matrix[i]) != cols {
			t.Fatalf("row %d cols: got %d want %d", i, len(matrix[i]), cols)
		}
	}
}
