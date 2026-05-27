# Authentication Guide

This document describes Sendio's dual authentication system.

---

## API Key Authentication

Used for programmatic access (CI/CD pipelines, backend services, scripts).

### How It Works

1. An API key is created and stored with:
   - **Key prefix** — Displayed in the UI (e.g., `sk_sendio_xxxx`)
   - **Key hash** — SHA-256 of `(raw_key + salt)`, stored in DB
   - **Rate limit** — Determined by subscription plan

2. The client sends the raw key in the `Authorization` header:
   ```
   Authorization: Bearer sk_sendio_a1b2c3d4e5f6...
   ```

3. The `AuthRequired` middleware:
   - Extracts the Bearer token
   - Hashes it with the configured salt
   - Looks up the hash in the database
   - Sets the API key ID in Fiber locals for downstream handlers

### Rotating API Keys

If a key is compromised, mark it as inactive in the database and create a new one:

```sql
UPDATE api_keys SET is_active = false, deleted_at = now() WHERE id = 'your-key-uuid';
```

### Best Practices

- Never commit API keys to version control
- Use environment variables for key storage in production
- Rotate keys periodically
- Use separate keys for separate environments (dev/staging/prod)

---

## Session Authentication (Web UI)

Used for browser-based access to the dashboard.

### How It Works

1. User clicks "Login with Google" or "Login with GitHub"
2. Browser redirects to OAuth provider
3. User grants permission
4. OAuth provider redirects back with a code
5. Server exchanges the code for tokens
6. Server creates a user (if new) and a 32-byte random session token
7. Token is hashed with `SESSION_SECRET` and stored in the database
8. Raw token is set as an `HttpOnly` cookie in the browser
9. Subsequent requests automatically include the cookie
10. `SessionAuth` middleware validates the session on every request

### Session Lifecycle

| Event | Action |
|-------|--------|
| Login | Create session (30-day expiry) |
| Request | Validate session, extend expiry |
| Logout | Delete session from DB, clear cookie |
| Expiry | Automatic cleanup via `DELETE expired sessions` |

### Cookie Settings

- `HttpOnly: true` — Not accessible via JavaScript (CSRF prevention)
- `Secure: true` — Only sent over HTTPS in production
- `SameSite: Lax` — Sent with same-site navigation
- `MaxAge` or `Expires` — 30-day session

---

## OAuth Providers

### Google

Scopes requested:
- `https://www.googleapis.com/auth/userinfo.email`
- `https://www.googleapis.com/auth/userinfo.profile`

Redirect URL: `/auth/callback/google`

### GitHub

Scopes requested:
- `user:email` — Access primary email
- `read:user` — Access public profile

Redirect URL: `/auth/callback/github`

Both providers handle email privacy — if a GitHub user has a hidden email, the API fetches the verified email list separately. Avatar URLs are stripped of the `=s...c` suffix for full-resolution images.

---

## Middleware Stack

When both authentication methods are in use, requests flow through:

```
Request
  │
  ▼
SessionAuth middleware
  │ (session cookie → validate → set user in locals)
  ▼
AuthRequired middleware
  │ (bearer token → validate → set apiKeyID in locals)
  ▼
Handler
```