MODULE := $(shell go list -m)

# Env loading from .env file
-include .env
export

GOOSE_DRIVER := postgres
GOOSE_DBSTRING := $(DB_URL)
GOOSE_MIGRATION_DIR := db/migration

.PHONY: init migrate-up migrate-down

init:
	# Remove all references to ${MODULE} in all config files
	grep -rl '\$${MODULE}' . --exclude=Makefile --exclude='*.go' | xargs perl -pi -e 's|\$${MODULE}|$(MODULE)|g'
	buf dep update
	buf generate .
	sqlc generate
	go mod tidy

db-migration-up:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up

db-migration-down:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" down

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go
