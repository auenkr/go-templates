MODULE := $(shell go list -m)
PLATFORM := linux/$(shell go env GOARCH)

# Env loading from .env file
-include .env
export

GOOSE_DRIVER := postgres
GOOSE_DBSTRING := $(DB_URL)
GOOSE_MIGRATION_DIR := db/migration

.PHONY: init generate run build build-image db-migration-status db-migration-up db-migration-down

init:
	# Remove all references to ${MODULE} in all config files
	find . -type f ! -name 'Makefile' ! -name '*.go' -exec sh -c \
		'sed '\''s|\$${MODULE}|$(MODULE)|g'\'' "$$1" > "$$1.tmp" && mv "$$1.tmp" "$$1"' _ {} \;
	buf dep update
	buf generate .
	sqlc generate
	go mod tidy

generate:
	sqlc generate
	buf generate .

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

build-image:
	docker build --platform ${PLATFORM} -f ./deploy/Dockerfile .

db-migration-up:
	goose up

db-migration-down:
	goose down

db-migration-status:
	goose status
