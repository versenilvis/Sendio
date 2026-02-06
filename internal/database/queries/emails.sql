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
