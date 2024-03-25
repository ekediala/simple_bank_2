package tx

import (
	"context"
	"database/sql"
	"fmt"

	database "github.com/ekediala/simple_bank_2/database/sqlc"
)

type TxManager struct {
	*database.Queries
	db *sql.DB
}

func New(db *sql.DB) *TxManager {
	return &TxManager{
		Queries: database.New(db),
		db:      db,
	}
}

func (txManager *TxManager) executeTransaction(ctx context.Context, fn func(*database.Queries) error) error {
	tx, err := txManager.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	queries := database.New(tx)

	err = fn(queries)

	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx error: %v; rollback error: %v", err, rbError)
		}
		return err
	}

	return tx.Commit()

}
