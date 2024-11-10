# Makefile

build:
	go build

test:
	go test ./...

coverage:
	go test -cover ./...

bench:
	go test -bench=. ./...

lint:
	golangci-lint run

run:
	go run main.go
