-- name: FindCustomerByID :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: FindCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1 LIMIT 1;

-- name: CreateCustomer :one
INSERT INTO customers (id, first_name, last_name, email, phone, date_of_birth, address_street, address_city, address_state, address_zip_code, address_country, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;

-- name: UpdateCustomer :exec
UPDATE customers
SET first_name = $2, last_name = $4, email = $3, phone = $5, date_of_birth = $6, address_street = $7, address_city = $8, address_state = $9, address_zip_code = $10, address_country = $11, status = $12
WHERE id = $1;
