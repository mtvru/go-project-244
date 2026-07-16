.PHONY: build run test test-coverage lint install

install:
	go mod download

build:
	go build -o bin/gendiff ./cmd/gendiff

run:
	go run ./cmd/gendiff

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run
