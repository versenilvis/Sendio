-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, name, avatar_url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $2, avatar_url = $3, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpsertOAuthAccount :exec
INSERT INTO oauth_accounts (user_id, provider, provider_id, access_token, refresh_token)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (provider, provider_id)
DO UPDATE SET
    access_token = EXCLUDED.access_token,
    refresh_token = EXCLUDED.refresh_token,
    updated_at = now();

-- name: CreateSession :exec
INSERT INTO sessions (user_id, token_hash, expires_at)
VALUES ($1, $2, $3);

-- name: GetSessionByTokenHash :one
SELECT s.*, u.email, u.name, u.avatar_url, u.is_active
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.token_hash = $1 AND s.expires_at > now()
LIMIT 1;

-- name: DeleteSessionByTokenHash :exec
DELETE FROM sessions WHERE token_hash = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at < now();
