# ---------- Build Stage ----------
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://proxy.golang.org,direct
ENV CGO_ENABLED=0

RUN apk add --no-cache git

# Copy go.mod files first (for caching)
COPY shared/go.mod shared/go.sum ./shared/
COPY product-service/go.mod product-service/go.sum ./product-service/

# Copy shared module
COPY shared/ ./shared/

# Set working directory
WORKDIR /app/product-service

# Cache dependencies
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy application source
COPY product-service/ .

# Build with cache
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -o main cmd/server/main.go

# ---------- Runtime Stage ----------
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/product-service/main .
EXPOSE 8082
CMD ["./main"]