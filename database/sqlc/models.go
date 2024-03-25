// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
)

type Account struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
}

type Entry struct {
	ID        int64        `json:"id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	AccountID int64        `json:"account_id"`
	// it can be negative or positive
	Amount int64 `json:"amount"`
}

type Transfer struct {
	ID            int64        `json:"id"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	FromAccountID int64        `json:"from_account_id"`
	ToAccountID   int64        `json:"to_account_id"`
	// only positive integers allowed
	Amount      int64 `json:"amount"`
	FromEntryID int64 `json:"from_entry_id"`
	ToEntryID   int64 `json:"to_entry_id"`
}
