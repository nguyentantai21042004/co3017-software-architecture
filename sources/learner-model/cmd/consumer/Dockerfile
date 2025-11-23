# syntax=docker/dockerfile:1

# =============================================================================
# Dockerfile for Learner Model Consumer
# Service: Event Consumer - Processes scoring events from RabbitMQ
# Port: None (background worker)
# Technology: Golang 1.23 + RabbitMQ + PostgreSQL
# =============================================================================

# =============================================================================
# STAGE 1: BUILDER
# Multi-platform build support for efficient compilation
# =============================================================================
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

WORKDIR /build

# Declare build platform arguments
ARG TARGETOS
ARG TARGETARCH

# Install build dependencies
RUN apk add --no-cache git tzdata ca-certificates

# Copy and download dependencies (cached layer)
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy entire project (context is learner-model/ directory)
COPY . .

# Build consumer binary with optimizations
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build \
    -ldflags="-s -w -X main.Version=1.0.0" \
    -o learner-consumer ./cmd/consumer

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
COPY --from=builder --chown=nonroot:nonroot /build/learner-consumer .

# Run as non-root user
USER nonroot:nonroot

# Health check metadata
LABEL org.opencontainers.image.title="Learner Model Consumer" \
      org.opencontainers.image.description="Background worker processing scoring events" \
      org.opencontainers.image.version="1.0.0"

ENTRYPOINT ["./learner-consumer"]
