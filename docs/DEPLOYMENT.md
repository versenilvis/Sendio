# Deployment Guide

This guide covers deploying Sendio for production use.

---

## Prerequisites

- Go 1.26+ for building the binary
- PostgreSQL 16+ (or managed service like Supabase, Neon, CockroachDB)
- Redis 7+ (or managed service like Redis Cloud, Upstash)
- A domain name for OAuth callback URLs

---

## Build

### Backend

```bash
# Build the API server
go build -o bin/server cmd/server/main.go

# Build the background worker
go build -o bin/worker cmd/worker/main.go
```

### Frontend

```bash
cd web
bun run build
```

The built frontend output goes into `web/build/` (SvelteKit Node adapter).

---

## Environment Variables

### Required

| Variable | Description |
|----------|-------------|
| `DATABASE_URL` | PostgreSQL connection string, e.g. `postgres://user:pass@host:5432/sendio?sslmode=require` |
| `API_KEY_SALT` | Random 32+ character string for hashing API keys |
| `SESSION_SECRET` | Random 32+ character string for signing sessions |
| `ENCRYPTION_KEY` | **Exactly 32 characters** for AES-256-GCM encryption |

Generate secure values:

```bash
openssl rand -hex 32   # for API_KEY_SALT, SESSION_SECRET
openssl rand -hex 16   # for ENCRYPTION_KEY (16 bytes = 32 hex chars)
```

### Optional

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `APP_ENV` | `development` | Set to `production` for production |
| `PUBLIC_SERVER_URL` | `http://localhost:8080` | Public URL for OAuth callbacks |
| `WEB_URL` | `http://localhost:5173` | Frontend URL |
| `GOOGLE_CLIENT_ID` | — | Google OAuth App Client ID |
| `GOOGLE_CLIENT_SECRET` | — | Google OAuth App Client Secret |
| `GITHUB_CLIENT_ID` | — | GitHub OAuth App Client ID |
| `GITHUB_CLIENT_SECRET` | — | GitHub OAuth App Client Secret |
| `REDIS_URL` | `redis://localhost:6379` | Redis connection URL |
| `SMTP_HOST` | `localhost` | Default SMTP relay host |
| `SMTP_PORT` | `1025` | Default SMTP relay port |
| `LOG_PATH` | `log/` | Log file directory |
| `DB_MAX_CONNS` | `25` | Max PostgreSQL connections |
| `DB_MIN_CONNS` | `5` | Min PostgreSQL connections |

---

## OAuth Setup

### Google OAuth

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Navigate to **APIs & Services → Credentials**
4. Create **OAuth Client ID** (Web application)
5. Add authorized redirect URI: `https://yourdomain.com/auth/callback/google`
6. Copy Client ID and Client Secret to your environment

### GitHub OAuth

1. Go to GitHub Settings → Developer settings → OAuth Apps
2. Create a new OAuth App
3. Set Homepage URL to your frontend URL
4. Set Authorization callback URL to `https://yourdomain.com/auth/callback/github`
5. Copy Client ID and Client Secret to your environment

---

## Database Setup

### Local with Docker

```bash
docker compose up -d postgres redis
```

### Managed Services

For production, use a managed PostgreSQL provider (Neon, Supabase, CockroachDB, etc.) and managed Redis (Redis Cloud, Upstash, etc.). Update `DATABASE_URL` and `REDIS_URL` accordingly.

### Running Migrations

Migrations run automatically on server startup. Alternatively, run them manually:

```bash
go run cmd/server/main.go migrate
```

Or using goose directly:

```bash
goose -dir migrations postgres "DATABASE_URL" up
```

---

## Running the Services

### Process Manager (Recommended)

Use systemd or a process manager like PM2, Supervisor, or systemd.

#### systemd (recommended)

**`/etc/systemd/system/sendio-server.service`**:

```ini
[Unit]
Description=Sendio API Server
After=network.target

[Service]
Type=simple
User=sendio
WorkingDirectory=/opt/sendio
ExecStart=/opt/sendio/bin/server
Environment=APP_ENV=production
EnvironmentFile=/opt/sendio/.env
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

**`/etc/systemd/system/sendio-worker.service`**:

```ini
[Unit]
Description=Sendio Background Worker
After=network.target

[Service]
Type=simple
User=sendio
WorkingDirectory=/opt/sendio
ExecStart=/opt/sendio/bin/worker
Environment=APP_ENV=production
EnvironmentFile=/opt/sendio/.env
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl enable sendio-server sendio-worker
sudo systemctl start sendio-server sendio-worker
```

### Docker Compose (Full Stack)

For a complete self-contained deployment:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: sendio
      POSTGRES_USER: sendio
      POSTGRES_PASSWORD: your_password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  server:
    build:
      context: .
      dockerfile: Dockerfile.server
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
    env_file:
      - .env
    depends_on:
      - postgres
      - redis
    restart: unless-stopped

  worker:
    build:
      context: .
      dockerfile: Dockerfile.worker
    env_file:
      - .env
    depends_on:
      - postgres
      - redis
    restart: unless-stopped

volumes:
  postgres_data:
  redis_data:
```

---

## Nginx Configuration

If using Nginx as a reverse proxy:

```nginx
server {
    listen 443 ssl;
    server_name api.sendio.example.com;

    ssl_certificate /etc/ssl/certs/your-cert.pem;
    ssl_certificate_key /etc/ssl/private/your-key.pem;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}
```

---

## SMTP Configuration

Sendio can relay email through your own SMTP server or a transactional email service (SendGrid, Amazon SES, Postmark, etc.).

Configure SMTP in the database via the SMTP configs table, or use the default SMTP settings via environment variables. SMTP passwords stored in the database are AES-256-GCM encrypted.

For production, it is recommended to configure:
- SPF records for your sending domain
- DKIM signing
- DMARC policy

---

## Security Checklist

- [ ] `DATABASE_URL` uses SSL mode (`sslmode=require`)
- [ ] `ENCRYPTION_KEY` is exactly 32 characters
- [ ] `API_KEY_SALT` and `SESSION_SECRET` are random and not shared
- [ ] OAuth credentials configured for all providers in use
- [ ] `APP_ENV=production` in production
- [ ] HTTPS configured on all public endpoints
- [ ] OAuth callback URLs use HTTPS
- [ ] Database credentials are not committed to version control
- [ ] Log files are not publicly accessible
- [ ] Redis is not exposed publicly (or use AUTH password)