-- name: CreateCustomerEvent :one
INSERT INTO events (id, context_id, event_origin, event_type, event_type_version, event_state, created_at, scheduled_at, retry, max_retry, event_data)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: FindCustomerEventByID :one
SELECT * FROM events 
WHERE id = $1 LIMIT 1;

