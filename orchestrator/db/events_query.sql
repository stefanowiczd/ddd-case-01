-- name: FindEvents :many
SELECT * FROM events
ORDER BY scheduled_at DESC;

-- name: FindProcessableEvents :many
SELECT * FROM events
WHERE event_state = 'ready'
ORDER BY scheduled_at DESC
LIMIT (sqlc.arg('limit'));

-- name: FindEventByID :one
SELECT * FROM events
WHERE id = $1;

-- name: FindEventsByOriginAndStatus :many
SELECT * FROM events
WHERE event_origin = $1 AND event_state = $2
ORDER BY scheduled_at DESC
LIMIT (sqlc.arg('limit'));

-- name: UpdateEventStartedAt :exec
UPDATE events
SET started_at = CURRENT_TIMESTAMP,
    event_state = 'processing'
WHERE id = $1;

-- name: UpdateEventRetry :exec
UPDATE events
SET retry = retry + 1,
    event_state = CASE
        WHEN retry + 1 >= max_retry THEN 'failed'
        ELSE 'ready'
    END,
    scheduled_at = CASE
        WHEN retry + 1 < max_retry THEN CURRENT_TIMESTAMP + (sqlc.arg('retry_interval') * INTERVAL '1 minute')
        ELSE scheduled_at
    END,
    completed_at = CASE
        WHEN retry + 1 >= max_retry THEN CURRENT_TIMESTAMP
    END
WHERE id = $1;

-- name: UpdateEventStart :exec
UPDATE events
SET started_at = CURRENT_TIMESTAMP,
    event_state = 'processing'
WHERE id = $1;


-- name: UpdateEventCompletion :exec
UPDATE events
SET completed_at = CURRENT_TIMESTAMP,
    event_state = 'completed'
WHERE id = $1;
