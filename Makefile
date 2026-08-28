include .env
export

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

COVER_PKGS := $$(go list ./internal/... | grep -v '/testutil' | paste -sd, -)
coverage:
	@echo "Testing and checking coverage..."
	@go test -coverprofile=coverage.out -coverpkg=$(COVER_PKGS) ./internal/...
	@go tool cover -html=coverage.out

check-env:
	@if [ -z "$$FLAGSMITH_API_KEY" ]; then \
		echo "⚠️ Variable FLAGSMITH_API_KEY is empty!"; \
		exit 1; \
	else \
		echo "✅ Variable FLAGSMITH_API_KEY is filled."; \
	fi

dev: vet check-env
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