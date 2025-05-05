-- name: FindEvents :many
SELECT * FROM events
ORDER BY scheduled_at DESC;

-- name: FindEventsByOrigin :many
SELECT * FROM events
WHERE event_origin = $1
ORDER BY scheduled_at DESC;

-- name: FindEventsByOriginAndStatus :many
SELECT * FROM events
WHERE event_origin = $1 AND event_state = $2
ORDER BY scheduled_at DESC;

-- name: SetEventState :exec
UPDATE events
SET event_state = $2
WHERE id = $1;
