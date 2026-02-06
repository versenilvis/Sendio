-- name: GetDefaultSMTPConfig :one
SELECT * FROM smtp_configs 
WHERE is_active = true AND is_default = true 
LIMIT 1;

-- name: GetSMTPConfigByID :one
SELECT * FROM smtp_configs WHERE id = $1;
