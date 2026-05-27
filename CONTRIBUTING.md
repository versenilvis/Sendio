# Contributing to Sendio

Thank you for your interest in contributing to Sendio. This guide covers everything you need to get started with development.

---

## Getting Started

### 1. Prerequisites

Install the required tools:

- **Go 1.26+** — Backend development
- **PostgreSQL 16+** — Database
- **Redis 7+** — Queue backend
- **Docker** — For infrastructure services
- **Bun** — Frontend package manager
- **sqlc** — Query generation (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)

### 2. Clone & Setup

```bash
git clone https://github.com/versenilvis/Sendio.git
cd Sendio

# Start infrastructure
docker compose up -d

# Copy environment file
cp .env.example .env
# Edit .env with your values
```

### 3. Install Frontend Dependencies

```bash
cd web
bun install
```

### 4. Generate Go Code

If you modify SQL files, regenerate the Go query layer:

```bash
sqlc generate
```

### 5. Run Server & Worker

```bash
# Terminal 1: API server
go run cmd/server/main.go

# Terminal 2: Background worker
go run cmd/worker/main.go
```

### 6. Run Frontend

```bash
cd web
bun run dev
```

---

## Project Structure

See [README.md](./README.md#architecture) for the overall architecture.

Key directories:

| Path | Purpose |
|------|---------|
| `cmd/server/` | HTTP API server entry point |
| `cmd/worker/` | Background job worker entry point |
| `internal/handlers/` | HTTP request handlers |
| `internal/repository/` | Database access layer (Repository pattern) |
| `internal/database/queries/` | SQL source files for sqlc |
| `internal/database/dbgen/` | Auto-generated Go code (do not edit) |
| `migrations/` | goose SQL migrations |
| `web/src/` | SvelteKit frontend |

---

## Code Conventions

### Go

- Run `gofmt` and `goimports` on all files before committing
- Wrap all errors with context using `fmt.Errorf("...: %w", err)`
- Use the Repository pattern for all database access
- Place SQL in `internal/database/queries/` and regenerate with `sqlc generate`

### Svelte / TypeScript

- Use Svelte 5 runes syntax (`$state`, `$props`, `$derived`) — not Svelte 4 options API
- Type all exported functions and component props explicitly
- No `console.log` in production code
- Run `bun run check` for type checking
- Run `bun run lint` for linting
- Run `bun run format` to format

---

## Testing

### Backend

```bash
go test ./...
go test -race ./...      # with race detection
go test -cover ./...     # with coverage
```

### Frontend

```bash
cd web
bun run test             # unit tests
bun run check            # TypeScript type checking
```

---

## Making Changes

1. **Fork** the repository
2. **Create afeature branch**: `git checkout -b feat/my-feature`
3. **Write tests first** (see TDD workflow in project rules)
4. **Implement** the feature
5. **Run checks**:
   ```bash
   go fmt ./... && go vet ./...
   # or for frontend
   bun run check && bun run lint
   ```
6. **Commit** using conventional commits format:
   ```
   feat: add audience management API
   fix: prevent duplicate email sends with idempotency key
   refactor: extract email rendering into separate package
   ```
7. **Push** and open a Pull Request

---

## Environment Variables

See [README.md](./README.md#configuration-reference) for the full list.

**Required** for any meaningful development:

```
DATABASE_URL=postgres://user:password@localhost:5432/sendio
API_KEY_SALT=your-random-32-char-salt-here
SESSION_SECRET=your-random-session-secret-here
ENCRYPTION_KEY=exactly-32-characters-long!!
```

---

## Adding New Database Queries

1. Write SQL in `internal/database/queries/<table>.sql`
2. Run `sqlc generate`
3. Implement the repository method in `internal/repository/`
4. Call the repository from the service layer

---

## Frontend Routing

All route pages are under `web/src/routes/`. Each route has a corresponding URL:

| Route | URL |
|-------|-----|
| `routes/dashboard/+page.svelte` | `/dashboard` |
| `routes/audiences/+page.svelte` | `/audiences` |
| `routes/broadcasts/+page.svelte` | `/broadcasts` |
| `routes/docs/api/+page.svelte` | `/docs/api` |

Always create both `+page.svelte` and its directory when adding routes.
