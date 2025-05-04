-- name: FindAccountByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: FindAccountsByCustomerID :many
SELECT * FROM accounts
WHERE customer_id = $1;

-- name: CreateAccount :one
INSERT INTO accounts (id, customer_id, account_number, balance, currency, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateAccountStatus :exec
UPDATE accounts
SET status = $2
WHERE id = $1;

-- name: DepositAccountMoney :exec
UPDATE accounts
SET balance = balance + $2
WHERE id = $1;

-- name: WithdrawAccountMoney :exec
UPDATE accounts
SET balance = balance - $2
WHERE id = $1;



