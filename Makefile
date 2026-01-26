.PHONY: up down run worker logs clean

up:
	@docker compose up -d

down:
	@docker compose down

run: # or just use "air"
	@go run cmd/server/main.go

worker w:
	@go run cmd/worker/main.go

logs:
	@docker compose logs -f

clean:
	@go clean
	@rm -rf bin/
