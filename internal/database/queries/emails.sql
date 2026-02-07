-- name: GetEmailByID :one
SELECT * FROM emails WHERE id = $1;

-- name: UpdateEmailStatus :exec
UPDATE emails SET status = $1 WHERE id = $2;

-- name: UpdateEmailStatusAtomic :execresult
-- This ensures idempotency: only update if current status allows it
UPDATE emails 
SET status = $1 
WHERE id = $2 AND status IN ('pending', 'queued');

-- name: UpdateEmailSent :exec
UPDATE emails SET status = $1, sent_at = NOW() WHERE id = $2;

-- name: UpdateEmailQueued :exec
UPDATE emails SET status = $1, queued_at = NOW() WHERE id = $2;
-- name: CreateEmail :one
INSERT INTO emails (
    api_key_id, idempotency_key, from_email, from_name, 
    to_email, to_name, subject, template_id, template_vars, 
    body_html, body_text, tags, metadata, status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: GetEmailByStatus :many
SELECT * FROM emails WHERE status = $1 LIMIT $2;