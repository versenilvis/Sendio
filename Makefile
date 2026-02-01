MAKEFLAGS += --no-print-directory

.PHONY: up down run worker logs clean web server migrate-up mu migrate-down md

up:
	@docker compose up -d $(s)

down:
	@docker compose down

run: # dont run this in production, it will cause logs to overlap
	@make server & make web

server s: # or just use "air"
	@go run ./cmd/server/main.go

web:
	@cd ./web && bun run dev

worker w:
	@go run cmd/worker/main.go

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

logs:
	@docker compose logs -f

migrate-up mu:
	@goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down md:
	@goose -dir migrations postgres "$(DATABASE_URL)" down

clean:
	@go clean
	@rm -rf bin/
