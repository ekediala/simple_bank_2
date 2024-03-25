// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entries.sql

package database

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (account_id, amount) VALUES($1, $2) RETURNING id, created_at, updated_at, account_id, amount
`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountID,
		&i.Amount,
	)
	return i, err
}

const deleteEntriesForAccount = `-- name: DeleteEntriesForAccount :exec
DELETE FROM entries WHERE account_id = $1
`

func (q *Queries) DeleteEntriesForAccount(ctx context.Context, accountID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntriesForAccount, accountID)
	return err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, entryID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, entryID)
	return err
}

const getEntriesForAccount = `-- name: GetEntriesForAccount :many
SELECT id, created_at, updated_at, account_id, amount FROM entries WHERE account_id = $3 LIMIT $1 OFFSET $2
`

type GetEntriesForAccountParams struct {
	Limit     int64 `json:"limit"`
	Offset    int64 `json:"offset"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) GetEntriesForAccount(ctx context.Context, arg GetEntriesForAccountParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesForAccount, arg.Limit, arg.Offset, arg.AccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccountID,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEntry = `-- name: GetEntry :one
SELECT id, created_at, updated_at, account_id, amount FROM entries WHERE id = $1
`

func (q *Queries) GetEntry(ctx context.Context, entryID int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, entryID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountID,
		&i.Amount,
	)
	return i, err
}

const getEntryForUpdate = `-- name: GetEntryForUpdate :one
SELECT id, created_at, updated_at, account_id, amount FROM entries WHERE id = $1 FOR NO KEY UPDATE
`

func (q *Queries) GetEntryForUpdate(ctx context.Context, entryID int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntryForUpdate, entryID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountID,
		&i.Amount,
	)
	return i, err
}

const updateEntry = `-- name: UpdateEntry :one
UPDATE entries SET amount = $1, updated_at = NOW() WHERE id = $2 RETURNING id, created_at, updated_at, account_id, amount
`

type UpdateEntryParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntry, arg.Amount, arg.ID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccountID,
		&i.Amount,
	)
	return i, err
}
