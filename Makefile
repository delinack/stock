LOCAL_BIN=$(CURDIR)/bin
NAME=stock

build:
			go build -o $(LOCAL_BIN)/$(NAME) cmd/app/main.go

# 1. Поднимает без ошибок приложение с помощью Docker контейнеров
# 2. Инфраструктура: база данных, миграции, данные для тестирования работы приложения
up:
			docker-compose up --build

run:		build
			$(LOCAL_BIN)/$(NAME)

install-goimports:
			GOBIN=$(LOCAL_BIN) go install github.com/pav5000/smartimports/cmd/smartimports@latest

install-linter:
			GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

install-goose:
			GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

install-gomock:
			GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0


bin-deps:	install-goimports install-linter install-goose install-gomock

format: 	install-goimports
			$(LOCAL_BIN)/smartimports .

linter:		install-linter
			$(LOCAL_BIN)/golangci-lint --config=.golangci.yml run ./...

mocks:		install-gomock
			$(LOCAL_BIN)/mockgen github.com/delinack/stock/internal/pkg/storage Storage > internal/pkg/mock/storage/storage_mock.go


#----------DB----------#
db:
			psql -U postgres -c "drop database if exists stock"
			psql -U postgres -c "create database stock"

goose-up:
			goose -allow-missing -dir migrations -table goose_db_version postgres "user=postgres password=qwerty host=localhost dbname=stock sslmode=disable" up

goose-down:
			goose -allow-missing -dir migrations -table goose_db_version postgres "user=postgres password=qwerty host=localhost dbname=stock sslmode=disable" down
#----------DB----------#

clean:
			rm -rf $(LOCAL_BIN)

.PHONY: install-goimports bin-deps format build run install-linter linter clean db goose-up goose-down install-goose