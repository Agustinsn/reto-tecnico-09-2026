# Matrix API - Go

REST API developed with Go and Fiber for QR factorization of rectangular matrices.

## Current milestone

- Fiber v3
- POST `/api/v1/qr`
- GET `/api/v1/health`
- Matrix validation
- Householder QR factorization
- Reduced QR output with `Q (m x k)` and `R (k x n)` where `k = min(m,n)`
- Unit tests validating `Q * R ~= A`
- Dockerfile using a multi-stage build

## Run locally

This project targets Go 1.27 and Fiber v3.

```bash
go mod tidy
go test ./...
go run ./cmd/server
```

## Example

```bash
curl -X POST http://localhost:8080/api/v1/qr \
  -H 'Content-Type: application/json' \
  -d '{"matrix":[[1,2],[3,4],[5,6]]}'
```

## Architecture

The QR calculation is isolated in `internal/services`, while the HTTP transport lives in `internal/handlers`. The next milestone adds an HTTP client in `internal/clients` to send `Q` and `R` to the Node.js statistics API.
