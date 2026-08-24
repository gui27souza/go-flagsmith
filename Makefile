.DEFAULT_GOAL := build

.PHONY: fmt vet test coverage dev build start clean

fmt:
	@echo "Formating..."
	@go fmt ./...

vet: fmt
	@echo "Vetting..."
	@go vet ./...

test:
	@echo "Testing..."
	@go test -v ./...

coverage:
	@echo "Testing and checking coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

dev: vet
	@echo "Running code..."
	@go run ./cmd/api

build: vet
	@echo "Building..."
	@mkdir -p bin
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/api ./cmd/api

start: build
	@echo "Running binary artifact..."
	@./bin/api

clean:
	@echo "Cleaning artifacts..."
	@rm -rf bin/
	@rm -f coverage.out