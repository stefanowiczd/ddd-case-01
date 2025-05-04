-- Create events table which is used to store all types of events
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    context_id UUID NOT NULL,
    event_origin VARCHAR(25) NOT NULL,
    event_type VARCHAR(25) NOT NULL,
    event_type_version VARCHAR(10) NOT NULL,
    event_state VARCHAR(25) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    scheduled_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    retry INT NOT NULL,
    max_retry INT NOT NULL,
    event_data JSONB NOT NULL
);

-- Create indexes for events table (draft)
CREATE INDEX IF NOT EXISTS idx_events_context_id ON events(context_id);
CREATE INDEX IF NOT EXISTS idx_events_event_origin ON events(event_origin);
CREATE INDEX IF NOT EXISTS idx_events_event_type ON events(event_type);
CREATE INDEX IF NOT EXISTS idx_events_scheduled_at ON events(scheduled_at);
