start-api: build run-api

start-worker: build run-worker

build:
	@echo ">> Building API..."
	@go build --race -o go-backend ./cmd
	@echo ">> Finished"

run-api:
	@./go-backend server

run-worker:
	@./go-backend worker

wire:
	@cd internal/app/init-module && wire

test:
	@go test -cover -race ./...

swag-init:
	@swag init -g cmd/main.go

start-components:
	@cd files/etc/docker && docker-compose --profile start-components up -d

start-redis:
	@cd files/etc/docker && docker-compose --profile start-redis up

start-postgres:
	@cd files/etc/docker && docker-compose --profile start-postgres up

stop-components:
	@cd files/etc/docker && docker-compose --profile start-components down

install-gomock:
	@echo "Installing Gomock"
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen@v1.6.0
	@echo "Finished installing gomock"