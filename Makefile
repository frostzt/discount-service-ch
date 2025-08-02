.PHONY: all fmt lint test

lint:
	@echo "> Running golangci-lint..."
	@golangci-lint run ./...

format:
	@echo "> Formatting with go fmt..."
	@go fmt ./...

test:
	@echo "> Running tests..."
	@go test -v ./...

all: fmt lint test