package tx

import (
	"context"

	database "github.com/ekediala/simple_bank_2/database/sqlc"
)

type TransferTxResult struct {
	FromTransfer database.Transfer `json:"from_transfer"`
	ToTransfer   database.Transfer `json:"to_transfer"`
	FromAccount  database.Account  `json:"from_account_id"`
	ToAccount    database.Account  `json:"to_account_id"`
	FromEntry    database.Entry    `json:"from_entry_id"`
	ToEntry      database.Entry    `json:"to_entry_id"`
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func updateAccountBalance(acc1ID, acc2ID int64, amount1, amount2 int64, queries *database.Queries, ctx context.Context) (acc1, acc2 database.Account, err error) {

	first, err := queries.GetAccountForUpdate(ctx, acc1ID)

	if err != nil {
		return
	}

	acc1, err = queries.UpdateAccountBalance(ctx, database.UpdateAccountBalanceParams{
		Amount: amount1,
		ID:     first.ID,
	})

	if err != nil {
		return
	}

	second, err := queries.GetAccountForUpdate(ctx, acc2ID)
	if err != nil {
		return
	}

	acc2, err = queries.UpdateAccountBalance(ctx, database.UpdateAccountBalanceParams{
		Amount: amount2,
		ID:     second.ID,
	})

	return acc1, acc2, err
}

func (txManager *TxManager) TransferTx(ctx context.Context, args TransferTxParams) (result TransferTxResult, err error) {
	err = txManager.executeTransaction(ctx, func(q *database.Queries) error {
		// create entry records
		result.FromEntry, err = q.CreateEntry(ctx, database.CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, database.CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})

		if err != nil {
			return err
		}

		// create transfer records
		result.FromTransfer, err = q.CreateTransfer(ctx, database.CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			ToEntryID:     result.ToEntry.ID,
			FromEntryID:   result.FromEntry.ID,
			Amount:        -args.Amount,
		})

		if err != nil {
			return err
		}

		result.ToTransfer, err = q.CreateTransfer(ctx, database.CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			ToEntryID:     result.ToEntry.ID,
			FromEntryID:   result.FromEntry.ID,
			Amount:        args.Amount,
		})
		// update account balances
		// ensure predictable order of updating accounts. remember we are updating both
		// concurrently, if there is no predictability we risk running into a deadlock
		if args.FromAccountID > args.ToAccountID {
			result.FromAccount, result.ToAccount, err = updateAccountBalance(args.FromAccountID, args.ToAccountID, -args.Amount, args.Amount, q, ctx)
			return err
		}

		result.ToAccount, result.FromAccount, err = updateAccountBalance(args.ToAccountID, args.FromAccountID, args.Amount, -args.Amount, q, ctx)
		return err
	})

	return result, err
}
