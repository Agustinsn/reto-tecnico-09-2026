export async function login(username: string, password: string): Promise<string> {
    const API_URL = import.meta.env.VITE_API_URL;
    const response = await fetch(`${API_URL}/api/v1/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error ?? "Error de autenticación");
  }

  return data.token;
}