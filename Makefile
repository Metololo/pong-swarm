.DEFAULT_GOAL := help

.PHONY: help build run test 

help:
	@echo "make build  Build the server"
	@echo "make run    Build and run the server"

build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

run: build
	./bin/server
