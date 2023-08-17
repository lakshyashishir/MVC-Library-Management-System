.PHONY: build run setup

all: setup build test run

setup:
	./scripts/config.sh

build:
	go mod vendor
	go mod tidy
	go build -o mvc ./cmd/main.go

test:
	go test ./pkg/models

run:
	./mvc
