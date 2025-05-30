INSERT INTO events (
    id,
    context_id,
    event_origin,
    event_type,
    event_state,
    event_type_version,
    created_at,
    scheduled_at,
    retry,
    max_retry,
    event_data
) VALUES 
(
  '00000000-0000-0000-0000-000000000000',
  '11111111-0000-0000-0000-000000000000',
  'customer',
  'customer.created',
  'ready',
  '0.0.0',
  '2025-05-06 00:00:00',
  '2025-05-06 00:00:00',
  0,
  3,
  '{
    "id": "00000000-0000-0000-0000-000000000000",
    "context_id": "11111111-0000-0000-0000-000000000000",
    "origin": "customer",
    "type": "customer.created",
    "type_version": "1.0.0",
    "state": "ready",
    "created_at": "2025-05-06T08:09:14.345074Z",
    "scheduled_at": "2025-05-06T08:09:14.345074Z",
    "started_at": "0001-01-01T00:00:00Z",
    "completed_at": "0001-01-01T00:00:00Z",
    "retry": 0,
    "max_retry": 3,
    "data": null,
    "firstName": "John",
    "lastName": "Doe",
    "phone": "1234567890",
    "email": "john.doe.3@example.com",
    "dateOfBirth": "1990-01-01",
    "address": {
      "street": "Street 1",
      "city": "Warsaw",
      "state": "Masovian",
      "postalCode": "00-000",
      "country": "Poland"
    }
  }'
),
(
  '00000000-0000-0000-0000-111111111111',
  '22222222-0000-0000-0000-000000000000',
  'customer',
  'customer.blocked',
  'ready',
  '0.0.0',
  '2025-05-06 00:00:00',
  '2025-05-06 00:00:00',
  2,
  3,
  '{
    "id": "00000000-0000-0000-0000-111111111111",
    "context_id": "22222222-0000-0000-0000-000000000000",
    "origin": "customer",
    "type": "customer.blocked",
    "type_version": "1.0.0",
    "state": "failed",
    "created_at": "2025-05-06T08:09:14.345074Z",
    "scheduled_at": "2025-05-06T08:09:14.345074Z",
    "started_at": "0001-01-01T00:00:00Z",
    "completed_at": "0001-01-01T00:00:00Z",
    "retry":2,
    "max_retry": 3,
    "data": null,
    "firstName": "John",
    "lastName": "Doe",
    "phone": "1234567890",
    "email": "john.doe.3@example.com",
    "dateOfBirth": "1990-01-01",
    "address": {
      "street": "Street 1",
      "city": "Warsaw",
      "state": "Masovian",
      "postalCode": "00-000",
      "country": "Poland"
    }
  }'
),
(
  '00000000-0000-0000-0000-222222222222',
  '33333333-0000-0000-0000-000000000000',
  'customer',
  'customer.unblocked',
  'failed',
  '0.0.0',
  '2025-05-06 00:00:00',
  '2025-05-06 00:00:00',
  3,
  3,
  '{
    "id": "00000000-0000-0000-0000-222222222222",
    "context_id": "33333333-0000-0000-0000-000000000000",
    "origin": "customer",
    "type": "customer.unblocked",
    "type_version": "1.0.0",
      "state": "failed",
    "created_at": "2025-05-06T08:09:14.345074Z",
    "scheduled_at": "2025-05-06T08:09:14.345074Z",
    "started_at": "0001-01-01T00:00:00Z",
    "completed_at": "0001-01-01T00:00:00Z",
    "retry": 3,
    "max_retry": 3,
    "data" : null,
    "firstName": "John",
    "lastName": "Doe",
    "phone": "1234567890",
    "email": "john.doe.3@example.com",
    "dateOfBirth": "1990-01-01",
    "address": {
      "street": "Street 1",
      "city": "Warsaw",
      "state": "Masovian",
      "postalCode": "00-000",
      "country": "Poland"
    }
  }'
);
