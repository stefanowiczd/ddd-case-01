INSERT INTO customers (
    id,
    first_name,
    last_name,
    email,
    date_of_birth,
    status
) VALUES (
    '00000000-0000-0000-0000-000000000000',
    'John',
    'Doe',
    'john.doe@example.com',
    '1990-01-01',
    'active'
);

---

INSERT INTO accounts (
    id,
    customer_id,
    account_number,
    balance,
    currency
) VALUES (
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000',
    '1234567890',
    1000.0,
    'USD'
);