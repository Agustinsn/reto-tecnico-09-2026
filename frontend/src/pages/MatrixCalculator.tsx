import { useState } from "react";
import { calculateQR } from "../services/qr";
import type { Matrix, QRResponse } from "../types/qr";

interface MatrixCalculatorProps {
  onLogout: () => void;
}

export function MatrixCalculator({ onLogout }: MatrixCalculatorProps) {
  const [rows, setRows] = useState(3);
  const [columns, setColumns] = useState(3);

  const [matrix, setMatrix] = useState<Matrix>([
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
  ]);

  const [result, setResult] = useState<QRResponse | null>(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  function createMatrix(newRows: number, newColumns: number) {
    const safeRows = Math.max(1, newRows);
    const safeColumns = Math.max(1, newColumns);

    const newMatrix = Array.from({ length: safeRows }, (_, i) =>
      Array.from({ length: safeColumns }, (_, j) => matrix[i]?.[j] ?? 0),
    );

    setRows(safeRows);
    setColumns(safeColumns);
    setMatrix(newMatrix);
    setResult(null);
  }

  function updateCell(row: number, column: number, value: string) {
    const newMatrix = matrix.map((currentRow, i) =>
      currentRow.map((cell, j) =>
        i === row && j === column ? Number(value) : cell,
      ),
    );

    setMatrix(newMatrix);
  }

  async function handleCalculate() {
    setError("");
    setResult(null);
    setLoading(true);

    try {
      const data = await calculateQR(matrix);
      setResult(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Error inesperado");
    } finally {
      setLoading(false);
    }
  }

  function handleLogout() {
    localStorage.removeItem("token");
    onLogout();
  }

  function renderMatrix(matrixToRender: Matrix) {
    return (
      <table className="simple-matrix">
        <tbody>
          {matrixToRender.map((row, i) => (
            <tr key={i}>
              {row.map((value, j) => (
                <td key={`${i}-${j}`}>{value.toFixed(5)}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    );
  }

  return (
    <main className="calculator-page">
      <header className="calculator-header">
        <div>
          <h1>Matrix QR Calculator</h1>
          <p>Calculadora de factorización QR</p>
        </div>

        <button onClick={handleLogout}>Cerrar sesión</button>
      </header>

      <section className="calculator-section">
        <h2>Matriz</h2>

        <div className="matrix-size">
          <label>
            Filas
            <input
              type="number"
              min="1"
              value={rows}
              onChange={(event) =>
                createMatrix(Number(event.target.value), columns)
              }
            />
          </label>

          <label>
            Columnas
            <input
              type="number"
              min="1"
              value={columns}
              onChange={(event) =>
                createMatrix(rows, Number(event.target.value))
              }
            />
          </label>
        </div>

        <table className="input-matrix">
          <tbody>
            {matrix.map((row, i) => (
              <tr key={i}>
                {row.map((value, j) => (
                  <td key={`${i}-${j}`}>
                    <input
                      type="number"
                      value={value}
                      onChange={(event) => updateCell(i, j, event.target.value)}
                    />
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>

        <button
          className="calculate-button"
          onClick={handleCalculate}
          disabled={loading}
        >
          {loading ? "Calculando..." : "Calcular"}
        </button>

        {error && <p className="calculator-error">{error}</p>}
      </section>

      {result && (
        <section className="results">
          <h2>Resultados</h2>

          <div className="result-block">
            <h3>Matriz Q</h3>
            {renderMatrix(result.q)}
          </div>

          <div className="result-block">
            <h3>Matriz R</h3>
            {renderMatrix(result.r)}
          </div>

          <div className="statistics">
            <h3>Estadísticas</h3>

            <h4>Q</h4>
            <p>Máximo: {result.statistics.q.max.toFixed(5)}</p>
            <p>Mínimo: {result.statistics.q.min.toFixed(5)}</p>
            <p>Promedio: {result.statistics.q.average.toFixed(5)}</p>
            <p>Suma: {result.statistics.q.sum.toFixed(5)}</p>
            <p>Diagonal: {result.statistics.q.isDiagonal ? "Sí" : "No"}</p>

            <h4>R</h4>
            <p>Máximo: {result.statistics.r.max.toFixed(5)}</p>
            <p>Mínimo: {result.statistics.r.min.toFixed(5)}</p>
            <p>Promedio: {result.statistics.r.average.toFixed(5)}</p>
            <p>Suma: {result.statistics.r.sum.toFixed(5)}</p>
            <p>Diagonal: {result.statistics.r.isDiagonal ? "Sí" : "No"}</p>

            <p>
              <strong>
                ¿Alguna matriz es diagonal?{" "}
                {result.statistics.anyDiagonal ? "Sí" : "No"}
              </strong>
            </p>
          </div>
        </section>
      )}
    </main>
  );
}
