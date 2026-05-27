# Sendio

**Modern email infrastructure for builders who care about delivery.**

Sendio is a self-hosted transactional email platform that gives you full control over your email infrastructure. Built with a Go backend and a SvelteKit frontend, it provides a clean API for sending transactional emails, managing SMTP configurations, and tracking delivery events — without depending on third-party services like SendGrid or Mailgun.

---

## Features

### Core

- **Transactional Email API** — Send emails via a simple REST API using API key authentication. Supports both raw HTML/text bodies and template-based rendering with variable substitution.
- **Template System** — Create reusable email templates with variable placeholders. Templates are stored in the database and rendered at send time.
- **SMTP Relay** — Configure your own SMTP server or use fallback SMTP configs. SMTP passwords are encrypted with AES-256-GCM before storage.
- **Idempotency** — Prevent duplicate sends by passing an idempotency key with each email request. The system ensures only one email is created per key.
- **Background Processing** — Email delivery is handled by an Asynq worker that pulls jobs from Redis. This keeps the API fast and allows for reliable retries on failure.

### Authentication

- **OAuth 2.0** — Sign in with Google or GitHub. OAuth tokens are stored and refreshed automatically. Web sessions use a SHA-256 hashed cookie with a 30-day expiry.
- **API Keys** — Full API key authentication for programmatic access. Keys are hashed with a salt before storage. Rate limiting is enforced per key.

### Delivery & Reliability

- **Status Tracking** — Emails move through states: `pending → queued → processing → sent → delivered`. Failed and bounced states are also tracked.
- **Retry Logic** — Failed email sends are retried up to 3 times with exponential backoff via Asynq.
- **Delivery Attempts** — Every SMTP attempt is logged with status, SMTP code/message, error details, and duration.
- **Event Logging** — Email lifecycle events (created, queued, sent, delivered, opened, clicked, bounced, complained) are recorded.

### Management & Quotas

- **Plan & Quota System** — Three built-in plans: Free (8,000 emails/month, 50/min rate), Premium ($19.99), Enterprise ($39.99). Quotas track sent count and reset on a billing cycle.
- **Audience Management** — Create audience segments and manage contacts for broadcast campaigns.
- **Broadcast Campaigns** — Send bulk email campaigns to audience segments with template support.

### Frontend

- **SvelteKit Dashboard** — Modern web UI built with Svelte 5 (runes syntax: `$state`, `$props`) and TailwindCSS v4.
- **Responsive Design** — Full responsive layout with navigation, hero sections, pricing, documentation, and more.

---

## Architecture

```
sendio/
├── cmd/
│   ├── server/          # Fiber HTTP server (health check + handlers)
│   └── worker/          # Asynq background job worker
├── internal/
│   ├── auth/            # API key verification
│   ├── config/           # Environment variable loading & validation
│   ├── crypto/           # AES-256-GCM encryption for SMTP passwords
│   ├── database/
│   │   ├── db.go        # PostgreSQL connection pool (pgx)
│   │   ├── redis.go     # Redis client
│   │   ├── queries/     # SQL queries (sqlc source)
│   │   └── dbgen/       # Generated Go code from sqlc
│   ├── email/           # Email service (create DB record → enqueue job)
│   ├── handlers/        # Fiber HTTP handlers (email, OAuth)
│   ├── logger/          # Zap structured logging with file rotation
│   ├── middleware/      # Auth middleware (API key + session)
│   ├── models/          # Domain models
│   ├── oauth/           # Google & GitHub OAuth providers + session manager
│   ├── queue/           # Asynq client (enqueue jobs)
│   ├── renderer/        # HTML template renderer (go template)
│   ├── repository/      # Repository pattern (Email, SMTP, Auth, User)
│   ├── utils/           # Utilities (env, close)
│   └── worker/          # Email processor (dequeue → SMTP → update status)
├── migrations/          # goose database migrations
├── web/                 # SvelteKit frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── assets/    # Static assets (avatar, etc.)
│   │   │   ├── components/
│   │   │   │   ├── navbar/  # NotchNav, etc.
│   │   │   │   └── hero/    # Hero, Features, Comparison
│   │   │   ├── config/     # App config constants
│   │   │   ├── constants/  # Terms, privacy policy text
│   │   │   ├── stores/     # Svelte 5 runes state (user store)
│   │   │   └── types/      # TypeScript types
│   │   └── routes/         # SvelteKit pages
│   └── package.json
├── docker-compose.yml   # PostgreSQL 16, Redis 7, Mailpit (local SMTP testing)
├── go.mod
├── go.sum
└── Makefile
```

### Data Flow

1. **API Request** → `AuthRequired` middleware validates Bearer token
2. **Handler** → `EmailHandler.SendEmail()` validates request
3. **Service** → `email.Service.SendEmail()` creates DB record → enqueues to Redis
4. **Worker** → Asynq dequeues `email:send` task → locks email record → sends via SMTP
5. **Worker** → Updates email status to `sent` or `failed` in DB

### Key Libraries

| Layer | Technology |
|-------|------------|
| HTTP Framework | [Fiber v3](https://gofiber.io/) |
| Database | PostgreSQL via [pgx](https://github.com/jackc/pgx) |
| Query Generation | [sqlc](https://sqlc.dev/) |
| Migrations | [goose](https://github.com/pressly/goose) |
| Queue | [Asynq](https://github.com/hibiken/asynq) + Redis |
| SMTP Client | [go-mail](https://github.com/wneessen/go-mail) |
| OAuth | [golang.org/x/oauth2](https://pkg.go.dev/golang.org/x/oauth2) |
| Logging | [Zap](https://go.uber.org/zap) |
| Frontend | [SvelteKit](https://kit.svelte.dev/) + Svelte 5 + TailwindCSS v4 |

---

## Quick Start

### Prerequisites

- Go 1.26+
- PostgreSQL 16+
- Redis 7+
- Bun (for frontend)

### 1. Clone & Configure

```bash
git clone https://github.com/versenilvis/Sendio.git
cd Sendio
```

Copy `.env.example` to `.env` and fill in the required values:

```bash
cp .env.example .env
```

Required environment variables:

| Variable | Description |
|----------|-------------|
| `DATABASE_URL` | PostgreSQL connection string |
| `API_KEY_SALT` | Salt for hashing API keys (min 32 chars) |
| `SESSION_SECRET` | Secret for signing session tokens (min 32 chars) |
| `ENCRYPTION_KEY` | 32-character key for AES-256-GCM encryption |
| `GOOGLE_CLIENT_ID` | Google OAuth Client ID *(optional)* |
| `GOOGLE_CLIENT_SECRET` | Google OAuth Client Secret *(optional)* |
| `GITHUB_CLIENT_ID` | GitHub OAuth App Client ID *(optional)* |
| `GITHUB_CLIENT_SECRET` | GitHub OAuth App Client Secret *(optional)* |

### 2. Start Infrastructure

```bash
docker compose up -d
```

This starts PostgreSQL (port 5432), Redis (port 6379), and Mailpit (SMTP at 1025, UI at 8025 for local testing).

### 3. Run Migrations

```bash
go run cmd/server/main.go
# Or run migrations standalone
```

The server automatically runs pending goose migrations on startup.

### 4. Start the Backend

```bash
# Terminal 1: API Server
go run cmd/server/main.go

# Terminal 2: Background Worker
go run cmd/worker/main.go
```

### 5. Start the Frontend

```bash
cd web
bun install
bun run dev
```

The frontend runs at `http://localhost:5173`. The API runs at `http://localhost:8080`.

---

## API Reference

### Authentication

All API requests require a Bearer token in the `Authorization` header:

```
Authorization: Bearer sk_xxxxxxxxxxxxxxxx
```

### Endpoints

#### `POST /api/v1/email/send`

Send a transactional email.

**Request Body:**

```json
{
  "from_email": "noreply@example.com",
  "from_name": "Sendio",
  "to_email": "user@example.com",
  "to_name": "John Doe",
  "subject": "Welcome to Sendio",
  "template_id": "550e8400-e29b-41d4-a716-446655440000",
  "template_vars": {
    "name": "John",
    "action_url": "https://example.com/verify"
  },
  "body_html": "<p>Hello world</p>",
  "body_text": "Hello world"
}
```

`from_email`, `to_email`, and `subject` are required. Either `template_id` + `template_vars` OR `body_html`/`body_text` must be provided.

**Response:**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "queued"
}
```

#### `GET /health`

Health check endpoint (no authentication required).

```json
{ "status": "ok" }
```

---

## Configuration Reference

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `APP_ENV` | `development` | Environment (`development` / `production`) |
| `PUBLIC_SERVER_URL` | `http://localhost:8080` | Server public URL (used for OAuth redirects) |
| `WEB_URL` | `http://localhost:5173` | Frontend URL |
| `DB_MAX_CONNS` | `25` | Maximum PostgreSQL connections |
| `DB_MIN_CONNS` | `5` | Minimum PostgreSQL connections |
| `REDIS_URL` | `redis://localhost:6379` | Redis connection URL |
| `SMTP_HOST` | `localhost` | Default SMTP host |
| `SMTP_PORT` | `1025` | Default SMTP port (Mailpit) |
| `LOG_PATH` | `log/` | Directory for log files |

---

## Database Schema

The schema is managed via two migration files:

1. **`000001_initial_schema.sql`** — Core tables: `plans`, `api_keys`, `api_key_usage`, `subscriptions`, `payment_transactions`, `email_templates`, `emails`, `smtp_configs`, `email_delivery_attempts`, `email_events`

2. **`000002_add_users_and_sessions.sql`** — OAuth and session tables: `users`, `oauth_accounts`, `sessions`

### Key Tables

- **`emails`** — Tracks every email with status, timestamps, and metadata
- **`api_keys`** — Stores hashed API keys with rate limits
- **`subscriptions`** — Links API keys to plans with quota tracking
- **`smtp_configs`** — Stores SMTP server configurations with encrypted passwords
- **`sessions`** — Stores SHA-256 hashed session tokens with expiry

---

## Email Status Flow

```
pending → queued → processing → sent → delivered
                              ↘ failed
                              ↘ bounced
                              ↘ cancelled
```

The worker uses `TryLockForProcessing` to ensure only one worker processes each email, providing idempotent delivery.

---

## Security

- **API Keys**: SHA-256 hashed with a per-deployment salt before storage
- **SMTP Passwords**: AES-256-GCM encrypted with a 32-character key + random 12-byte nonce
- **Sessions**: 32-byte random token, SHA-256 hashed with session secret before DB storage
- **All secrets**: Loaded from environment variables, never hardcoded

---

## Development

See [CONTRIBUTING.md](./CONTRIBUTING.md) for setup instructions and development guidelines.

See [docs/](./docs/) for detailed documentation on architecture, API, and deployment.