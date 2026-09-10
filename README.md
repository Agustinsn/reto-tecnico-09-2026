# Reto Técnico - APIs de Procesamiento de Matrices

Proyecto desarrollado como parte de un reto técnico para implementar dos APIs que trabajan de manera conjunta para procesar matrices numéricas.

La primera API está desarrollada en Go y se encarga de recibir una matriz y realizar su factorización QR.

La segunda API está desarrollada en Node.js y se encarga de calcular estadísticas sobre las matrices resultantes.

## ¿Cómo funciona?

El flujo principal es el siguiente:

1. El cliente envía una matriz a la API desarrollada en Go.
2. Go realiza la factorización QR de la matriz.
3. Go envía las matrices `Q` y `R` a la API desarrollada en Node.js.
4. Node.js calcula:
   - Valor máximo
   - Valor mínimo
   - Promedio
   - Suma total
   - Si la matriz es diagonal

5. Go devuelve al cliente el resultado de la factorización junto con las estadísticas.

```text
Cliente
   |
   v
Go API
   |
   | Q y R
   v
Node.js API
   |
   v
Estadísticas
   |
   v
Respuesta
```

## Tecnologías utilizadas

### API Go

- Go
- Fiber
- HTTP/REST
- Docker

### API Node.js

- Node.js
- TypeScript
- Express
- Vitest
- Docker

### Infraestructura

- Docker
- Docker Compose

## Estructura del proyecto

```text
reto-tecnico-fs/
│
├── api-go/
│   ├── cmd/
│   ├── internal/
│   │   ├── clients/
│   │   ├── domain/
│   │   ├── handlers/
│   │   ├── services/
│   │   └── validators/
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── api-node/
│   ├── src/
│   │   ├── controllers/
│   │   ├── routes/
│   │   ├── services/
│   │   ├── types/
│   │   └── validators/
│   ├── Dockerfile
│   ├── package.json
│   └── tsconfig.json
│
├── docker-compose.yml
└── README.md
```

## Requisitos

Para ejecutar el proyecto localmente se necesita:

- Go 1.27 o superior
- Node.js 20
- npm
- Docker Desktop

## Ejecución con Docker Compose

La forma más sencilla de ejecutar todo el proyecto es mediante Docker Compose.

Desde la raíz del proyecto:

```powershell
docker compose up --build
```

Esto inicia las dos APIs:

- Go API: `http://localhost:8080`
- Node.js API: `http://localhost:3000`

Para detener los servicios:

```powershell
docker compose down
```

## Endpoints

### Go API

Health check:

```http
GET /api/v1/health
```

Factorización QR:

```http
POST /api/v1/qr
```

Ejemplo de solicitud:

```json
{
  "matrix": [
    [1, 2],
    [3, 4],
    [5, 6]
  ]
}
```

La respuesta contiene las matrices `Q` y `R`, además de las estadísticas calculadas por la API de Node.js.

### Node.js API

Health check:

```http
GET /api/v1/health
```

Estadísticas:

```http
POST /api/v1/statistics
```

Ejemplo:

```json
{
  "q": [
    [1, 0],
    [0, 2]
  ],
  "r": [
    [3, 4],
    [0, 5]
  ]
}
```

## Pruebas

La API de Go cuenta con pruebas para la factorización QR y las validaciones de matrices.

Para ejecutarlas:

```powershell
cd api-go
go test ./...
```

La API de Node.js utiliza Vitest para probar el cálculo de estadísticas.

Para ejecutar las pruebas:

```powershell
cd api-node
npm test
```

También se puede verificar que el proyecto TypeScript compile correctamente:

```powershell
npm run build
```

## Factorización QR

La factorización QR se implementó utilizando reflexiones de Householder.

La implementación permite trabajar con matrices:

- Rectangulares altas
- Cuadradas
- Rectangulares anchas

Para matrices rectangulares se utiliza la factorización QR reducida.

La elección de Householder permite obtener una implementación estable para este tipo de cálculo.

## Validaciones

Las APIs validan que las matrices:

- No estén vacías.
- Tengan al menos una columna.
- Sean rectangulares.
- Contengan valores numéricos válidos.

Para determinar si una matriz es diagonal se considera una pequeña tolerancia numérica debido a las aproximaciones propias de los cálculos con números de punto flotante.

## Comunicación entre servicios

La comunicación entre Go y Node.js se realiza mediante HTTP.

Cuando se ejecuta con Docker Compose, Go utiliza el nombre del servicio de Node.js para comunicarse con él:

```text
http://api-node:3000
```

La URL puede configurarse mediante la variable de entorno:

```text
NODE_API_URL
```

## Estado del proyecto

Actualmente el proyecto cuenta con:

- API de procesamiento QR en Go.
- API de estadísticas en Node.js.
- Comunicación HTTP entre ambas APIs.
- Validación de matrices.
- Pruebas automatizadas.
- Dockerfiles para ambos servicios.
- Docker Compose para ejecutar el proyecto completo.

Como siguiente etapa se contempla el despliegue de las APIs en AWS.
