import type { Matrix, QRResponse } from "../types/qr";

const API_URL = "https://matrix-api-go.onrender.com";

export async function calculateQR(matrix: Matrix): Promise<QRResponse> {
  const token = localStorage.getItem("token");

  if (!token) {
    throw new Error("Sesión no válida");
  }

  const response = await fetch(`${API_URL}/api/v1/qr`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      matrix,
    }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error ?? "Error al calcular QR");
  }

  return data;
}