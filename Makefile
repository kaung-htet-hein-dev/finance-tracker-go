.PHONY: build run clean test docker-build docker-run

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Tidy dependencies
tidy:
	go mod tidy

# Build Docker image
docker-build:
	docker build -t finance-tracker .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Stop Docker Compose
docker-stop:
	docker-compose down

# Development setup
dev-setup: tidy build
	@echo "Development environment ready!"

# Help
help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  clean       - Clean build artifacts"
	@echo "  test        - Run tests"
	@echo "  fmt         - Format code"
	@echo "  tidy        - Tidy dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run  - Run with Docker Compose"
	@echo "  docker-stop - Stop Docker Compose"
	@echo "  dev-setup   - Setup development environment"