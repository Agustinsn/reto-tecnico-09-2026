import { Request, Response } from "express";
import { calculateStatistics } from "../services/statistics.service.js";
import { validateMatrix } from "../validators/matrix.validator.js";

interface StatisticsRequest {
  q: unknown;
  r: unknown;
}

export function calculateMatrixStatistics(
  req: Request,
  res: Response,
): void {
  const body = req.body as StatisticsRequest;

  if (!validateMatrix(body.q)) {
    res.status(400).json({
      error: "Invalid matrix q",
    });

    return;
  }

  if (!validateMatrix(body.r)) {
    res.status(400).json({
      error: "Invalid matrix r",
    });

    return;
  }

  const qStatistics = calculateStatistics(body.q);
  const rStatistics = calculateStatistics(body.r);

  res.json({
    q: qStatistics,
    r: rStatistics,
    anyDiagonal: qStatistics.isDiagonal || rStatistics.isDiagonal,
  });
}