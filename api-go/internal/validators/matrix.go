package validators

import (
	"fmt"

	"github.com/Agustinsn/reto-tecnico-09-2026/api-go/internal/domain"
)

// Validate returns a client-facing error when the matrix does not satisfy
// the API input contract.
func Validate(matrix domain.Matrix) error {
	if len(matrix) == 0 {
		return fmt.Errorf("matrix must contain at least one row")
	}

	if len(matrix[0]) == 0 {
		return fmt.Errorf("matrix must contain at least one column")
	}

	columns := len(matrix[0])
	for i, row := range matrix {
		if len(row) != columns {
			return fmt.Errorf("matrix is not rectangular: row %d has %d columns, expected %d", i, len(row), columns)
		}
	}

	return nil
}
