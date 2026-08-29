include .env
export

export CGO_ENABLED := 0

# OS Detection
ifeq ($(OS),Windows_NT)
	BINARY := bin\api.exe
	CLEAN_CMD := if exist bin rmdir /s /q bin && if exist coverage.out del /f /q coverage.out
else
	BINARY := ./bin/api
	CLEAN_CMD := rm -rf bin/ coverage.out
endif

.DEFAULT_GOAL := build

.PHONY: fmt vet test coverage check-env dev build start clean

fmt:
	@echo Formatting...
	@go fmt ./...

vet: fmt
	@echo Vetting...
	@go vet ./...

test:
	@echo Testing...
	@go test -v ./...

ALL_PKGS := $(shell go list ./internal/...)
FILTERED_PKGS := $(filter-out %/testutil, $(ALL_PKGS))
COMMA := ,
EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
COVER_PKGS := $(subst $(SPACE),$(COMMA),$(FILTERED_PKGS))

coverage:
	@echo Testing and checking coverage...
	@go test -coverprofile=coverage.out -coverpkg=$(COVER_PKGS) ./internal/...
	@go tool cover -html=coverage.out

check-env:
	@$(if $(strip $(FLAGSMITH_API_KEY)),echo ✅ Variable FLAGSMITH_API_KEY is filled.,$(error ⚠️ Variable FLAGSMITH_API_KEY is empty!))

dev: vet check-env
	@echo Running code...
	@go run ./cmd/api

build: vet
	@echo Building...
	@go build -ldflags="-s -w" -o $(BINARY) ./cmd/api

start: build
	@echo Running binary artifact...
	@$(BINARY)

clean:
	@echo Cleaning artifacts...
	@$(CLEAN_CMD)