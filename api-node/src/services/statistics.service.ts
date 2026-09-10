import { Matrix } from "../types/matrix.js";

export interface MatrixStatistics {
  max: number;
  min: number;
  average: number;
  sum: number;
  isDiagonal: boolean;
}

const DIAGONAL_EPSILON = 1e-10;

export function calculateStatistics(matrix: Matrix): MatrixStatistics {
  let sum = 0;
  let min = matrix[0][0];
  let max = matrix[0][0];

  for (const row of matrix) {
    for (const value of row) {
      sum += value;

      if (value < min) {
        min = value;
      }

      if (value > max) {
        max = value;
      }
    }
  }

  const totalElements = matrix.reduce(
    (total, row) => total + row.length,
    0,
  );

  return {
    max,
    min,
    average: sum / totalElements,
    sum,
    isDiagonal: isDiagonalMatrix(matrix),
  };
}

function isDiagonalMatrix(matrix: Matrix): boolean {
  const rowCount = matrix.length;
  const columnCount = matrix[0].length;

  // A diagonal matrix must be square.
  if (rowCount !== columnCount) {
    return false;
  }

  for (let row = 0; row < rowCount; row++) {
    for (let column = 0; column < columnCount; column++) {
      if (
        row !== column &&
        Math.abs(matrix[row][column]) > DIAGONAL_EPSILON
      ) {
        return false;
      }
    }
  }

  return true;
}