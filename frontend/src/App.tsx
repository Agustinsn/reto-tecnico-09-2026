import { useState } from "react";
import { Login } from "./pages/Login";
import "./App.css";
import { MatrixCalculator } from "./pages/MatrixCalculator";

function App() {
  const [token, setToken] = useState(() => localStorage.getItem("token") ?? "");

  if (!token) {
    return <Login onLogin={setToken} />;
  }

  return <MatrixCalculator onLogout={() => setToken("")} />;
}

export default App;
