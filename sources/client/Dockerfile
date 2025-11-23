# syntax=docker/dockerfile:1

# =============================================================================
# Dockerfile for Client Application
# Service: Frontend Client - Next.js web application
# Port: 3000
# Technology: Node.js 20 + Next.js 15 + React 19
# =============================================================================

# =============================================================================
# STAGE 1: DEPENDENCIES - Install dependencies
# =============================================================================
FROM node:20-alpine AS deps

WORKDIR /app

# Install dependencies for better compatibility
RUN apk add --no-cache libc6-compat

# Copy package files
COPY package.json package-lock.json* ./

# Install dependencies with legacy-peer-deps for React 19
RUN --mount=type=cache,target=/root/.npm \
    npm ci --legacy-peer-deps

# =============================================================================
# STAGE 2: BUILDER - Build the Next.js application
# =============================================================================
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependencies from deps stage
COPY --from=deps /app/node_modules ./node_modules

# Copy source code
COPY . .

# Set production environment for build
ENV NEXT_TELEMETRY_DISABLED=1 \
    NODE_ENV=production

# Build Next.js application
RUN npm run build

# =============================================================================
# STAGE 3: RUNNER - Production runtime
# =============================================================================
FROM node:20-alpine AS runner

WORKDIR /app

# Create non-root user
RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nextjs

# Set production environment
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000

# Copy necessary files from builder
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

# Set correct permissions
RUN chown -R nextjs:nodejs /app

# Switch to non-root user
USER nextjs

# Expose port
EXPOSE 3000

# Health check metadata
LABEL org.opencontainers.image.title="ITS Client Application" \
      org.opencontainers.image.description="Intelligent Tutoring System web frontend" \
      org.opencontainers.image.version="0.1.0"

# Start Next.js server
CMD ["node", "server.js"]
