-- name: FindEvents :many
SELECT * FROM events
ORDER BY scheduled_at DESC;

-- name: FindEventsByOrigin :many
SELECT * FROM events
WHERE event_origin = $1
ORDER BY scheduled_at DESC;

-- name: FindEventsByOriginAndType :many
SELECT * FROM events
WHERE event_origin = $1 AND event_type = $2
ORDER BY scheduled_at DESC;

-- name: FindEventsByOriginAndTypeAndState :many
SELECT * FROM events
WHERE event_origin = $1 AND event_type = $2 AND event_state = $3
ORDER BY scheduled_at DESC;
