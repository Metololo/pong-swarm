.DEFAULT_GOAL := help

.PHONY: help build run test 

help:
	@echo "  make build  Build the server"
	@echo "  make run    Build and run the server"
	@echo "  make fmt             - Format all code (gofmt)"
	@echo "  make lint            - Run linter on the whole project"
	@echo "  make fix             - Auto-fix fmt and lint issues (Global)"
build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

run: build
	./bin/Server

fmt:
	gofmt -s -w .

lint:
	golangci-lint run ./...

fix:
	gofmt -s -w .
	golangci-lint run --fix