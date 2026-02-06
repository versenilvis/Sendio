-- name: GetAPIKeyByHash :one
SELECT * FROM api_keys WHERE key_hash = $1 AND is_active = true LIMIT 1;

-- name: GetPlanByID :one
SELECT * FROM plans WHERE id = $1;

-- name: GetSubscriptionByApiKeyID :one
SELECT * FROM subscriptions WHERE api_key_id = $1 AND status = 'active' LIMIT 1;
