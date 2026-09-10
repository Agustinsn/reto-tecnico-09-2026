import { Matrix } from "../types/matrix.js";

export function validateMatrix(matrix: unknown): matrix is Matrix {
  if (!Array.isArray(matrix) || matrix.length === 0) {
    return false;
  }

  if (!Array.isArray(matrix[0]) || matrix[0].length === 0) {
    return false;
  }

  const columnCount = matrix[0].length;

  return matrix.every(
    (row) =>
      Array.isArray(row) &&
      row.length === columnCount &&
      row.every((value) => typeof value === "number" && Number.isFinite(value)),
  );
}