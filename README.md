# Reto Técnico - Calculadora QR

Aplicación web que permite ingresar una matriz y obtener su factorización QR, junto con algunos datos sobre las matrices resultantes.

## Tecnologías

- Go
- Node.js
- React + TypeScript
- Docker
- JWT para el inicio de sesión

## Estructura

```text
reto-tecnico-fs/
├── api-go/       # Cálculo QR y autenticación
├── api-node/     # Cálculo de estadísticas
├── frontend/     # Aplicación web
└── docker-compose.yml
```

## Uso local

### 1. API Node.js

```bash
cd api-node
npm install
npm run dev
```

Se ejecuta en `http://localhost:3000`.

### 2. API Go

Configurar las variables:

```text
AUTH_USERNAME=admin
AUTH_PASSWORD=admin123
JWT_SECRET=reto-tecnico-fs
NODE_API_URL=http://localhost:3000
```

Luego ejecutar:

```bash
cd api-go
go run ./cmd/server
```

Se ejecuta en `http://localhost:8080`.

### 3. Frontend

```bash
cd frontend
npm install
npm run dev
```

Se ejecuta en `http://localhost:5173`.

## Docker

Para ejecutar las APIs con Docker:

```bash
docker compose up --build
```

## Acceso

El proyecto está publicado en Render:

**Frontend:**
https://matrix-qr-frontend.onrender.com/

**API Go:**
https://matrix-api-go.onrender.com

**API Node.js:**
https://matrix-api-node.onrender.com

### Credenciales de prueba

```text
Usuario: admin
Contraseña: admin123
```

## Funcionalidades

- Inicio de sesión.
- Ingreso de matrices de diferentes tamaños.
- Cálculo de Q y R.
- Cálculo de máximo, mínimo, promedio y suma.
- Comprobación de matrices diagonales.
- Cierre de sesión.

## Pruebas

API Go:

```bash
go test ./... -count=1
```

API Node.js:

```bash
npm test
```
