export type Matrix = number[][];

export interface MatrixStatistics {
  max: number;
  min: number;
  average: number;
  sum: number;
  isDiagonal: boolean;
}

export interface QRStatistics {
  q: MatrixStatistics;
  r: MatrixStatistics;
  anyDiagonal: boolean;
}

export interface QRResponse {
  q: Matrix;
  r: Matrix;
  statistics: QRStatistics;
}