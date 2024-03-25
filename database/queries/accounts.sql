-- name: CreateAccount :one
INSERT INTO accounts(owner, currency) VALUES($1, $2) RETURNING *;

-- name: GetAccount :one
SELECT id, owner, balance, updated_at, created_at, currency FROM accounts WHERE id = sqlc.arg(account_id);

-- name: GetAccountForUpdate :one
SELECT id, owner, balance, updated_at, created_at, currency FROM accounts WHERE id = sqlc.arg(account_id) FOR NO KEY UPDATE;

-- name: UpdateAccountOwner :one
UPDATE accounts SET owner = $1, updated_at = NOW() WHERE id = $2 RETURNING *;

-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = balance + sqlc.arg(amount), updated_at = NOW() WHERE id = $2 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = sqlc.arg(account_id);
