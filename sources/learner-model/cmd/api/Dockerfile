# syntax=docker/dockerfile:1

# =============================================================================
# Dockerfile for Learner Model Service API
# Service: Learner Model - Student mastery tracking and skill management
# Port: 8083
# Technology: Golang 1.25.4 + Gin + PostgreSQL + MinIO + RabbitMQ
# =============================================================================

# =============================================================================
# STAGE 1: BUILDER
# Multi-platform build support for efficient compilation
# =============================================================================
FROM --platform=$BUILDPLATFORM golang:1.25.4-alpine AS builder

WORKDIR /build

# Declare build platform arguments
ARG TARGETOS
ARG TARGETARCH

# Install build dependencies
RUN apk add --no-cache git tzdata ca-certificates

# Copy and download dependencies (cached layer)
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && \
    go install github.com/swaggo/swag/cmd/swag@v1.16.6

# Copy entire project (context is learner-model/ directory)
COPY . .

# Generate Swagger documentation
RUN swag init -g cmd/api/main.go -o docs

# Build binary with optimizations
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build \
    -ldflags="-s -w -X main.Version=1.0.0" \
    -o learner-api ./cmd/api

# =============================================================================
# STAGE 2: RUNTIME - Distroless for security and minimal size
# =============================================================================
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy timezone and CA certificates from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set timezone
ENV TZ=Asia/Ho_Chi_Minh

# Copy binary
COPY --from=builder --chown=nonroot:nonroot /build/learner-api .

# Copy swagger docs if needed
COPY --from=builder --chown=nonroot:nonroot /build/docs ./docs

# Run as non-root user
USER nonroot:nonroot

# Expose service port
EXPOSE 8083

# Health check metadata
LABEL org.opencontainers.image.title="Learner Model Service API" \
      org.opencontainers.image.description="Student mastery tracking and learning analytics" \
      org.opencontainers.image.version="1.0.0"

ENTRYPOINT ["./learner-api"]