# Content Service Makefile
# Spring Boot 3.5.6, Java 17

.PHONY: help run-api swagger build clean test install

# Default target
.DEFAULT_GOAL := help

# Variables
JAR_FILE=target/content-service-0.0.1-SNAPSHOT.jar
MAIN_CLASS=co3017.microservices.content_service.ContentServiceApplication
SWAGGER_OUTPUT=docs/swagger.json
SWAGGER_YAML=docs/swagger.yaml
PORT=8081

# Help target
help:
	@echo "Content Service - Available targets:"
	@echo ""
	@echo "Development:"
	@echo "  run-api          - Run the API service (Spring Boot)"
	@echo "  swagger          - Generate Swagger/OpenAPI documentation"
	@echo "  build            - Build the application (Maven package)"
	@echo "  install          - Install dependencies and build"
	@echo ""
	@echo "Testing:"
	@echo "  test             - Run all tests"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo ""
	@echo "Utilities:"
	@echo "  clean            - Clean build artifacts"
	@echo "  run-jar          - Run the application from JAR file"
	@echo ""
	@echo "Examples:"
	@echo "  make run-api     - Start the service on port $(PORT)"
	@echo "  make swagger     - Generate OpenAPI spec files"
	@echo "  make build       - Build JAR file"

# Install dependencies and build
install:
	@echo "Installing dependencies and building..."
	@mvn clean install -DskipTests

# Build the application
build:
	@echo "Building Content Service..."
	@mvn clean package -DskipTests
	@echo "Build complete: $(JAR_FILE)"

# Run the API service (main entry point)
run-api:
	@echo "Starting Content Service API..."
	@echo "Swagger UI will be available at: http://localhost:$(PORT)/content/swagger/index.html"
	@echo "API will be available at: http://localhost:$(PORT)"
	@mvn spring-boot:run -Dmaven.test.skip=true

# Generate Swagger/OpenAPI documentation
swagger:
	@echo "Generating Swagger/OpenAPI documentation..."
	@mkdir -p docs
	@echo "Starting service to generate OpenAPI spec..."
	@echo "Note: SpringDoc OpenAPI generates spec at runtime"
	@echo "Swagger UI: http://localhost:$(PORT)/content/swagger/index.html"
	@echo "OpenAPI JSON: http://localhost:$(PORT)/v3/api-docs"
	@echo "OpenAPI YAML: http://localhost:$(PORT)/v3/api-docs.yaml"
	@echo ""
	@echo "To download the spec file, run the service and use:"
	@echo "  curl http://localhost:$(PORT)/v3/api-docs > $(SWAGGER_OUTPUT)"
	@echo "  curl http://localhost:$(PORT)/v3/api-docs.yaml > $(SWAGGER_YAML)"
	@echo ""
	@echo "Or use: make swagger-download (requires service to be running)"

# Download Swagger spec (requires service to be running)
swagger-download:
	@echo "Downloading OpenAPI spec from running service..."
	@mkdir -p docs
	@curl -s http://localhost:$(PORT)/v3/api-docs > $(SWAGGER_OUTPUT) || (echo "Error: Service not running. Start with 'make run-api' first" && exit 1)
	@curl -s http://localhost:$(PORT)/v3/api-docs.yaml > $(SWAGGER_YAML) || (echo "Error: Service not running. Start with 'make run-api' first" && exit 1)
	@echo "Swagger spec downloaded:"
	@echo "  - JSON: $(SWAGGER_OUTPUT)"
	@echo "  - YAML: $(SWAGGER_YAML)"

# Run from JAR file (requires build first)
run-jar: build
	@echo "Running Content Service from JAR..."
	@echo "Swagger UI: http://localhost:$(PORT)/content/swagger/index.html"
	@java -jar $(JAR_FILE)

# Run tests
test:
	@echo "Running tests..."
	@mvn test

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@mvn test jacoco:report
	@echo "Coverage report generated in target/site/jacoco/index.html"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@mvn clean
	@rm -rf docs/swagger.json docs/swagger.yaml
	@echo "Clean complete"

