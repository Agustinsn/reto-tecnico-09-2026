import { describe, expect, it } from "vitest";
import { calculateStatistics } from "./statistics.service.js";

describe("calculateStatistics", () => {
  it("should calculate max, min, average and sum", () => {
    const matrix = [
      [1, 2],
      [3, 4],
    ];

    const result = calculateStatistics(matrix);

    expect(result.max).toBe(4);
    expect(result.min).toBe(1);
    expect(result.sum).toBe(10);
    expect(result.average).toBe(2.5);
  });

  it("should detect a diagonal matrix", () => {
    const matrix = [
      [1, 0],
      [0, 2],
    ];

    const result = calculateStatistics(matrix);

    expect(result.isDiagonal).toBe(true);
  });

  it("should detect a non-diagonal matrix", () => {
    const matrix = [
      [1, 2],
      [0, 3],
    ];

    const result = calculateStatistics(matrix);

    expect(result.isDiagonal).toBe(false);
  });

  it("should consider a matrix with very small off-diagonal values as diagonal", () => {
    const matrix = [
      [1, 1e-12],
      [1e-12, 2],
    ];

    const result = calculateStatistics(matrix);

    expect(result.isDiagonal).toBe(true);
  });

  it("should consider a rectangular matrix as non-diagonal", () => {
    const matrix = [
      [1, 2, 3],
      [4, 5, 6],
    ];

    const result = calculateStatistics(matrix);

    expect(result.isDiagonal).toBe(false);
  });
});