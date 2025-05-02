-- name: CreateAccountEvent :one
INSERT INTO account_events (account_id, event_type, event_type_version, event_state, scheduled_at, retry, max_retry, event_data)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: FindAccountEventByID :one
SELECT * FROM account_events
WHERE id = $1
ORDER BY created_at DESC;

-- name: FindAccountEventsByAccountID :many
SELECT * FROM account_events
WHERE account_id = $1
ORDER BY created_at DESC;

-- name: FindAccountEventsTheNewestFirst :many
SELECT * FROM account_events
ORDER BY created_at DESC;

-- name: FindAccountEventsByAccountIDAndEventType :many
SELECT * FROM account_events
WHERE account_id = $1 AND event_type = $2
ORDER BY created_at DESC;

-- name: FindAccountEventsByDateRange :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2
ORDER BY created_at DESC;

-- name: FindAccountEventsByDateRangeAndAccountID :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2 AND account_id = $3
ORDER BY created_at DESC;

-- name: FindAccountEventsByDateRangeAndAccountIDAndEventType :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2 AND account_id = $3 AND event_type = $4
ORDER BY created_at DESC;
