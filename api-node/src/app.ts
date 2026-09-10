import express from "express";
import statisticsRoutes from "./routes/statistics.routes.js";

const app = express();

const PORT = Number(process.env.PORT) || 3000;

app.use(express.json());

app.get("/api/v1/health", (_req, res) => {
  res.json({
    status: "ok",
  });
});

app.use("/api/v1", statisticsRoutes);

app.listen(PORT, () => {
  console.log(`Node API listening on :${PORT}`);
});