// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: accounts.sql

package database

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts(owner, currency) VALUES($1, $2) RETURNING id, owner, created_at, updated_at, balance, currency
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Balance,
		&i.Currency,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, accountID)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, updated_at, created_at, currency FROM accounts WHERE id = $1
`

type GetAccountRow struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	CreatedAt sql.NullTime `json:"created_at"`
	Currency  string       `json:"currency"`
}

func (q *Queries) GetAccount(ctx context.Context, accountID int64) (GetAccountRow, error) {
	row := q.db.QueryRowContext(ctx, getAccount, accountID)
	var i GetAccountRow
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Currency,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, balance, updated_at, created_at, currency FROM accounts WHERE id = $1 FOR NO KEY UPDATE
`

type GetAccountForUpdateRow struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	CreatedAt sql.NullTime `json:"created_at"`
	Currency  string       `json:"currency"`
}

func (q *Queries) GetAccountForUpdate(ctx context.Context, accountID int64) (GetAccountForUpdateRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, accountID)
	var i GetAccountForUpdateRow
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Currency,
	)
	return i, err
}

const updateAccountBalance = `-- name: UpdateAccountBalance :one
UPDATE accounts SET balance = balance + $1, updated_at = NOW() WHERE id = $2 RETURNING id, owner, created_at, updated_at, balance, currency
`

type UpdateAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBalance, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Balance,
		&i.Currency,
	)
	return i, err
}

const updateAccountOwner = `-- name: UpdateAccountOwner :one
UPDATE accounts SET owner = $1, updated_at = NOW() WHERE id = $2 RETURNING id, owner, created_at, updated_at, balance, currency
`

type UpdateAccountOwnerParams struct {
	Owner string `json:"owner"`
	ID    int64  `json:"id"`
}

func (q *Queries) UpdateAccountOwner(ctx context.Context, arg UpdateAccountOwnerParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountOwner, arg.Owner, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Balance,
		&i.Currency,
	)
	return i, err
}
