MAKEFLAGS += --no-print-directory

.PHONY: up down run worker logs clean web server

up:
	@docker compose up -d

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

logs:
	@docker compose logs -f

clean:
	@go clean
	@rm -rf bin/
