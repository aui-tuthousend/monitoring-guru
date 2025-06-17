# Stage 1: Builder
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Install swag & air
RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/air-verse/air@latest

# Copy go mod files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source
COPY . .

# Generate Swagger docs
RUN swag init

# Build binary (opsional jika kamu ingin prebuild juga)
RUN go build -o app main.go

# Stage 2: Final image (untuk development, not scratch)
FROM golang:1.24.4-alpine AS dev

WORKDIR /app

# Install runtime dependencies (if any)
RUN apk add --no-cache bash

# Copy all source code
COPY --from=builder /app /app

# Copy air binary
COPY --from=builder /go/bin/air /usr/bin/air

# Expose the port used by Fiber
EXPOSE 8080

# Run with Air
CMD ["air", "-c", ".air.toml"]
