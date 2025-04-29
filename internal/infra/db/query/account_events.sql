-- name: CreateAccountEvent :one
INSERT INTO account_events (account_id, event_type, event_data)
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindAccountEventsByAccountID :many
SELECT * FROM account_events
WHERE account_id = $1;

-- name: FindAccountEventsTheNewestFirst :many
SELECT * FROM account_events
ORDER BY created_at DESC;

-- name: FindAccountEventsByAccountIDAndEventType :many
SELECT * FROM account_events
WHERE account_id = $1 AND event_type = $2;

-- name: FindAccountEventsByDateRange :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2;

-- name: FindAccountEventsByDateRangeAndAccountID :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2 AND account_id = $3;


-- name: FindAccountEventsByDateRangeAndAccountIDAndEventType :many
SELECT * FROM account_events
WHERE created_at BETWEEN $1 AND $2 AND account_id = $3 AND event_type = $4;

