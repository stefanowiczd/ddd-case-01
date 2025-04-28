--- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetCustomerAccounts :many
SELECT * FROM accounts
WHERE customer_id = $1;

-- name: CreateAccountStatus :one
INSERT INTO accounts (customer_id, account_number, balance, currency, status)
VALUES ($1, $2, $3, $4, $5)
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



