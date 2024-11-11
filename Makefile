# Makefile

run:
	go run main.go

build:
	go build

prod:
	go build -o dist/ -ldflags "-s -w" -trimpath

test:
	go test ./...

coverage:
	go test -cover ./...

bench:
	go test -bench=. ./...

lint:
	golangci-lint run

docker-build:
	docker build -t go-http-server-router .

docker-run:
	docker run -p 8080:8080 go-http-server-router

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down
