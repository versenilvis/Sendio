# API Reference

Base URL: `http://localhost:8080` (development)

---

## Authentication

Sendio supports two authentication methods depending on your use case.

### API Key (Programmatic Access)

For programmatic access to the email sending API.

```http
Authorization: Bearer sk_your_api_key_here
```

API keys are created per application and have configurable rate limits. The key prefix (`sk_`) is stored in the database for identification, but the full key is hashed with SHA-256 + salt.

### Session Cookie (Web UI)

For browser-based access to the dashboard. The session cookie is set automatically after completing an OAuth flow. It is `HttpOnly`, `Secure` (in production), and expires after 30 days.

---

## Endpoints

### Health Check

#### `GET /health`

Returns the health status of the API server. No authentication required.

**Response:**

```json
{
  "status": "ok"
}
```

---

### Email

#### `POST /api/v1/email/send`

Send a transactional email. The email is queued immediately and delivered asynchronously by the background worker.

**Authentication:** API Key (Bearer token)

**Request Body:**

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `from_email` | string | Yes | Sender email address |
| `from_name` | string | No | Sender display name |
| `to_email` | string | Yes | Recipient email address |
| `to_name` | string | No | Recipient display name |
| `subject` | string | Yes | Email subject line |
| `template_id` | string (UUID) | No* | Use a stored email template |
| `template_vars` | object | No* | Variables for template rendering |
| `body_html` | string | No* | Raw HTML email body |
| `body_text` | string | No | Plain text email body |

*Either `template_id` + `template_vars` OR `body_html`/`body_text` must be provided.

**Example — Using raw HTML:**

```bash
curl -X POST http://localhost:8080/api/v1/email/send \
  -H "Authorization: Bearer sk_your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "from_email": "noreply@sendio.dev",
    "from_name": "Sendio",
    "to_email": "user@example.com",
    "to_name": "Jane Doe",
    "subject": "Welcome to Sendio",
    "body_html": "<h1>Hello Jane!</h1><p>Welcome to our platform.</p>",
    "body_text": "Hello Jane! Welcome to our platform."
  }'
```

**Example — Using a template:**

```bash
curl -X POST http://localhost:8080/api/v1/email/send \
  -H "Authorization: Bearer sk_your_api_key" \
  -H "Content-Type: application/json" \
  -d '{
    "from_email": "noreply@sendio.dev",
    "to_email": "user@example.com",
    "subject": "Welcome, {{name}}!",
    "template_id": "550e8400-e29b-41d4-a716-446655440000",
    "template_vars": {
      "name": "Jane",
      "verify_url": "https://example.com/verify/abc123"
    }
  }'
```

**Response (HTTP 202 Accepted):**

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "queued"
}
```

**Error Responses:**

| Status | Body | Cause |
|--------|------|-------|
| `400` | `{"error": "to_email and subject are required"}` | Missing required fields |
| `400` | `{"error": "Invalid request body"}` | Malformed JSON |
| `401` | `{"error": "..."}` | Invalid or missing API key |
| `500` | `{"error": "..."}` | Internal error (DB, queue, etc.) |

---

### OAuth (Web UI)

#### `GET /auth/login/google`

Redirects to Google's OAuth consent screen.

#### `GET /auth/login/github`

Redirects to GitHub's OAuth authorization page.

#### `GET /auth/callback/google`

OAuth callback for Google. Exchanges the code for tokens, creates/updates the user, creates a session, and sets a session cookie. On success, closes the popup window via JavaScript.

#### `GET /auth/callback/github`

Same flow as Google callback.

#### `GET /auth/me`

Returns the currently authenticated user's information.

**Response:**

```json
{
  "email": "user@example.com",
  "name": "Jane Doe",
  "avatar_url": "https://avatars.githubusercontent.com/u/1234567",
  "is_active": true,
  "created_at": "2026-01-15T10:30:00Z"
}
```

#### `POST /auth/logout`

Invalidates the current session and clears the cookie.

**Response:**

```json
{
  "message": "Logged out successfully"
}
```

---

## Email Status Codes

| Status | Description |
|--------|-------------|
| `pending` | Email created, waiting to be picked up by worker |
| `queued` | Email enqueued to Asynq/Redis for processing |
| `processing` | Worker has claimed the email and is sending it |
| `sent` | SMTP send completed successfully |
| `delivered` | Recipient mail server confirmed delivery |
| `failed` | All send attempts exhausted (max 3 retries) |
| `bounced` | Recipient mail server rejected the message |
| `cancelled` | Manually cancelled before sending |

---

## Rate Limiting

Rate limits are enforced per API key based on the associated plan. The current rate limit for your key is included in plan metadata.

| Plan | Monthly Quota | Rate Limit |
|------|--------------|------------|
| Free | 8,000 emails | 50/min |
| Premium | Unlimited | configurable |
| Enterprise | Unlimited | configurable |

When a quota is exhausted, the API returns `429 Too Many Requests`.

---

## Webhooks & Events

Email lifecycle events are tracked in the `email_events` table:

| Event | Trigger |
|-------|---------|
| `created` | Email record inserted into DB |
| `queued` | Task enqueued to Asynq |
| `sent` | SMTP send succeeded |
| `delivered` | Recipient server acknowledged receipt |
| `opened` | Tracking pixel fired (future) |
| `clicked` | Link clicked (future) |
| `bounced` | Recipient server rejected message |
| `complained` | User marked as spam |
| `failed` | All retry attempts exhausted |

---

## Idempotency

To prevent duplicate sends, pass an `idempotency_key` field with each email request:

```json
{
  "idempotency_key": "order-12345-confirmation",
  "from_email": "noreply@example.com",
  "to_email": "user@example.com",
  "subject": "Order Confirmation #12345",
  "body_html": "<p>Your order is confirmed.</p>"
}
```

If the same idempotency key is used within a short window, the original email ID is returned without creating a duplicate.