# Architecture

This document describes the system architecture of Sendio.

---

## Overview

Sendio follows a layered architecture with clear separation of concerns:

```
┌─────────────────────────────────────────────┐
│              SvelteKit Frontend              │
│          (web/src - port 5173)              │
└──────────────────┬──────────────────────────┘
                   │ HTTP (session cookie)
┌──────────────────▼──────────────────────────┐
│              Fiber HTTP API                  │
│         (cmd/server - port 8080)             │
│  ┌─────────────┬─────────────┬────────────┐  │
│  │ Middleware  │  Handlers   │   Models   │  │
│  └─────────────┴─────────────┴────────────┘  │
└────────┬────────────────────────┬────────────┘
         │                        │
┌────────▼────────┐    ┌─────────▼────────────┐
│  email.Service  │    │  repository.Registry  │
└────────┬────────┘    └─────────┬────────────┘
         │                       │
┌────────▼────────┐    ┌─────────▼────────────┐
│  queue.Client  │    │     PostgreSQL        │
└────────┬────────┘    └─────────────────────┘
         │
┌────────▼────────┐
│     Redis        │
└────────┬────────┘
         │
┌────────▼────────┐
│ Asynq Worker    │
│ (cmd/worker)    │
└─────────────────┘
```

---

## Backend (Go)

### HTTP Layer (`internal/handlers/`)

HTTP handlers receive requests, validate input, delegate to service layer, and return responses. Each handler is focused on a single responsibility:

- **`email.go`** — Handles email send requests. Parses the `SendEmailRequest`, extracts the API key ID from middleware context, and calls `email.Service.SendEmail()`.
- **`oauth.go`** — Handles OAuth 2.0 flows: login redirects, callbacks, user info, and logout.

### Service Layer (`internal/email/`)

The `email.Service` is the core business logic:

1. Validates the request parameters
2. Creates an `email` record in PostgreSQL with status `pending`
3. Enqueues an `email:send` task to Redis via Asynq
4. Updates the email status to `queued`

The service is decoupled from the HTTP layer — it receives a Fiber context for request-scoped values.

### Authentication (`internal/auth/`, `internal/middleware/`)

Two authentication mechanisms coexist:

| Method | Use Case | Mechanism |
|--------|----------|-----------|
| **API Key** | Programmatic / machine access | `Authorization: Bearer sk_xxx` header. Key is SHA-256 hashed with salt, looked up in DB. |
| **Session** | Web UI / browser access | Cookie-based. 32-byte random token, hashed and stored in DB. 30-day expiry. |

Both mechanisms set values in Fiber's locals context for downstream handlers.

### Repository Pattern (`internal/repository/`)

All database access goes through a `Repository.Registry` that aggregates sub-repositories:

```
Registry
├── Email: email CRUD, idempotent locking
├── SMTP: SMTP config lookup
├── Auth: API key verification
└── User: user CRUD, session management
```

The `TryLockForProcessing` method in `EmailRepository` implements atomic idempotent consumer locking — it uses an atomic UPDATE to change status from `pending`/`queued` to `processing`, returning `true` only if the update affected a row. This ensures only one worker processes each email.

### Queue (`internal/queue/`)

Asynq client wraps task creation. The `EmailSendPayload` contains only the email UUID — the worker fetches all needed data from the DB after dequeuing. Tasks are configured with a maximum of 3 retries.

### Worker (`internal/worker/`)

The `Processor` handles `email:send` tasks:

1. **Lock** — Calls `TryLockForProcessing` to claim the email
2. **Fetch** — Loads the email record and SMTP config from DB
3. **Render** — If a template ID was provided, renders the HTML template with variables
4. **Build** — Constructs the SMTP message with From/To/Subject/Body
5. **Send** — Dials the SMTP server and sends the message
6. **Update** — Sets status to `sent` or `failed`

The worker also handles TLS configuration (NoTLS, TLS optional, SSL) and decrypts SMTP passwords on-the-fly using AES-256-GCM.

### Database (`internal/database/`)

- **Connection**: `pgx/v5/pgxpool` for connection pooling. Configured via `DATABASE_URL`.
- **Migrations**: `goose/v3` runs pending migrations on server startup.
- **Queries**: SQL files in `queries/` are compiled by `sqlc` into type-safe Go in `dbgen/`.
- **Models**: Two model layers — raw DB-generated structs (`dbgen/`) and clean domain models (`models/`). The repository maps between them.

### Security (`internal/crypto/`)

- **AES-256-GCM**: Used to encrypt SMTP passwords before storing in the database. Each encrypted value includes a random 12-byte nonce and is Base64-encoded.
- **API Key Hashing**: SHA-256 hash of (API key + salt) for lookup without storing raw keys.
- **Session Tokens**: SHA-256 hash of (raw token + session secret) stored in DB.

---

## Frontend (SvelteKit)

### Tech Stack

- **SvelteKit 5** — Full-stack framework with SSR/SSG
- **Svelte 5** — Uses the new runes API (`$state`, `$props`, `$derived`)
- **TailwindCSS v4** — Utility-first CSS via `@tailwindcss/vite`
- **Bun** — Package manager
- **Playwright** — E2E testing

### Routing

SvelteKit file-based routing under `web/src/routes/`. The `+layout.svelte` wraps all pages with global styles.

### State Management

Svelte 5 runes (`$state`) in the `stores/` directory. The `user.svelte.ts` file manages the current user state with usage statistics.

### Components

Components are organized by feature:
- `components/navbar/` — `NotchNav` and related navigation components
- `components/hero/` — Landing page hero section, features, comparison
- UI components are placed at the root of `components/`

### Environment

The `PUBLIC_SERVER_URL` environment variable is exposed to the frontend for API calls.

---

## Infrastructure

### Docker Compose

Three services:

| Service | Image | Ports | Purpose |
|---------|-------|-------|---------|
| `postgres` | `postgres:16-alpine` | 5432 | Primary data store |
| `redis` | `redis:7-alpine` | 6379 | Asynq queue backend |
| `mailpit` | `axllent/mailpit` | 1025 (SMTP), 8025 (UI) | Local SMTP testing / email preview |

### Logging

Zap structured logger writing to both stdout and `log/app.log` with log rotation (1MB max, 5 backups, 5 days). In development mode, logs are output to console with color encoding and stack traces enabled at DEBUG level.

---

## Email Lifecycle

```
1. Client POST /api/v1/email/send
         │
2. AuthRequired middleware validates Bearer token
         │
3. EmailHandler.SendEmail() validates input
         │
4. email.Service.SendEmail()
         ├── Creates email record (status=pending)
         ├── Enqueues email:send task to Redis
         └── Updates status=queued
         │
5. Asynq Worker dequeues task
         │
6. Processor.ProcessTask()
         ├── TryLockForProcessing (atomic status=processing)
         ├── Loads email + SMTP config from DB
         ├── Renders template (if template_id provided)
         ├── Builds SMTP message
         └── Sends via go-mail client
         │
7. Status updated to sent/failed
```

---

## Session Management

Web sessions use a cookie-based flow:

1. User clicks "Login with Google/GitHub"
2. OAuth redirect to provider → callback → session token generated
3. Session token (raw) set as `HttpOnly` cookie
4. Each request: `SessionAuth` middleware hashes token + secret, looks up in DB
5. User data attached to Fiber locals
6. Logout: token deleted from DB, cookie cleared

---

## Build & Run

See [DEPLOYMENT.md](./DEPLOYMENT.md) for production deployment instructions.
