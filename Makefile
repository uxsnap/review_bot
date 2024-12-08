include .env

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

run:
	nodemon --watch './internal/**/*.go' --signal SIGTERM --exec 'go' run ./cmd/main.go

down:
	docker-compose down
	docker volume rm fresh_market_shop_postgres_volume

restart-migrations:
	docker-compose restart migrator

build:
	docker-compose up -d