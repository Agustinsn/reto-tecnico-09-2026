FROM golang:1.27-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/matrix-api ./cmd/server

FROM alpine:3.22
RUN addgroup -S app && adduser -S app -G app
USER app
COPY --from=builder /bin/matrix-api /bin/matrix-api
EXPOSE 8080
ENTRYPOINT ["/bin/matrix-api"]
