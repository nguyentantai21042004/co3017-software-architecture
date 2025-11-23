# syntax=docker/dockerfile:1

# =============================================================================
# Dockerfile for Content Service
# Service: Content Service - Question and learning material management
# Port: 8081
# Technology: Java 17 + Spring Boot 3.5.6 + Maven + PostgreSQL
# =============================================================================

# =============================================================================
# STAGE 1: BUILDER - Build the Spring Boot application
# =============================================================================
FROM eclipse-temurin:17-jdk-alpine AS builder

WORKDIR /build

# Copy Maven wrapper and pom.xml first for better caching
COPY .mvn/ .mvn/
COPY mvnw pom.xml ./

# Download dependencies (cached layer if pom.xml unchanged)
RUN --mount=type=cache,target=/root/.m2 \
    ./mvnw dependency:go-offline -B || true

# Copy source code
COPY src/ src/

# Build the application
RUN --mount=type=cache,target=/root/.m2 \
    ./mvnw clean package -DskipTests -B

# Extract JAR layers for better Docker caching
RUN mkdir -p target/extracted && \
    java -Djarmode=layertools -jar target/content-service-*.jar extract --destination target/extracted

# =============================================================================
# STAGE 2: RUNTIME - Optimized JRE with Spring Boot layers
# =============================================================================
FROM eclipse-temurin:17-jre-alpine

WORKDIR /app

# Install curl for healthcheck
RUN apk add --no-cache curl

# Create non-root user
RUN addgroup -S spring && adduser -S spring -G spring
USER spring:spring

# Copy Spring Boot layers (most stable layers first for better caching)
COPY --from=builder --chown=spring:spring /build/target/extracted/dependencies/ ./
COPY --from=builder --chown=spring:spring /build/target/extracted/spring-boot-loader/ ./
COPY --from=builder --chown=spring:spring /build/target/extracted/snapshot-dependencies/ ./
COPY --from=builder --chown=spring:spring /build/target/extracted/application/ ./

# Set environment variables
ENV JAVA_OPTS="-Xms256m -Xmx512m -XX:+UseG1GC -XX:MaxGCPauseMillis=200" \
    TZ=Asia/Ho_Chi_Minh

# Expose service port
EXPOSE 8081

# Health check metadata
LABEL org.opencontainers.image.title="Content Service" \
      org.opencontainers.image.description="Question and content management service" \
      org.opencontainers.image.version="0.0.1-SNAPSHOT"

# Health check
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD curl -f http://localhost:8081/health || exit 1

# Run the application using Spring Boot's launcher
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS org.springframework.boot.loader.launch.JarLauncher"]