-- name: CreateEntry :one
INSERT INTO entries (account_id, amount) VALUES($1, $2) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries WHERE id = sqlc.arg(entry_id);

-- name: GetEntryForUpdate :one
SELECT * FROM entries WHERE id = sqlc.arg(entry_id) FOR NO KEY UPDATE;

-- name: GetEntriesForAccount :many
SELECT * FROM entries WHERE account_id = sqlc.arg(account_id) LIMIT $1 OFFSET $2;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = sqlc.arg(entry_id);

-- name: DeleteEntriesForAccount :exec
DELETE FROM entries WHERE account_id = sqlc.arg(account_id);

-- name: UpdateEntry :one
UPDATE entries SET amount = $1, updated_at = NOW() WHERE id = $2 RETURNING *;