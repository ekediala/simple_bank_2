-- name: CreateTransfer :one
INSERT INTO transfers(from_account_id, to_account_id, from_entry_id, to_entry_id, amount)
VALUES($1,$2,$3,$4,$5) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers WHERE id = sqlc.arg(transfer_id);

-- name: GetTransferForUpdate :one
SELECT * FROM transfers WHERE id = sqlc.arg(transfer_id) FOR NO KEY UPDATE;

-- name: UpdateTransferAmount :one
UPDATE transfers SET amount = sqlc.arg(transfer_amount), updated_at = NOW() WHERE id = sqlc.arg(transfer_id) RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = sqlc.arg(transfer_id);