import { Router } from "express";
import { calculateMatrixStatistics } from "../controllers/statistics.controller.js";

const router = Router();

router.post("/statistics", calculateMatrixStatistics);

export default router;