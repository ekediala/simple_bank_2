package tx_test

import (
	"context"
	"testing"

	database "github.com/ekediala/simple_bank_2/database/sqlc"
	"github.com/ekediala/simple_bank_2/database/tx"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

type TransferTxResponse struct {
	result    tx.TransferTxResult
	sender    database.Account
	recipient database.Account
	err       error
}

func createRandomAccount(t *testing.T) database.Account {
	fake := faker.New()
	owner := fake.Person().Name()

	data := database.CreateAccountParams{
		Owner:    owner,
		Currency: "USD",
	}

	account, err := testTx.Queries.CreateAccount(context.Background(), database.CreateAccountParams{
		Owner:    data.Owner,
		Currency: data.Currency,
	})

	require.NoError(t, err)
	require.NotEmpty(t, account)

	account, err = testTx.Queries.UpdateAccountBalance(context.Background(), database.UpdateAccountBalanceParams{
		Amount: 1000,
		ID:     account.ID,
	})

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func TestTransferTx(t *testing.T) {
	channel := make(chan TransferTxResponse)
	n := 10
	amount := int64(10)

	sender := createRandomAccount(t)
	recipient := createRandomAccount(t)

	for i := 0; i < n; i++ {
		go func() {
			result, err := testTx.TransferTx(context.Background(), tx.TransferTxParams{
				FromAccountID: sender.ID,
				ToAccountID:   recipient.ID,
				Amount:        amount,
			})
			channel <- TransferTxResponse{result: result, sender: sender, recipient: recipient, err: err}
		}()
	}

	// drain channel
	for i := 0; i < n; i++ {
		result := <-channel
		require.NoError(t, result.err)
		// ensure the expected data is all there
		require.NotEmpty(t, result.result.FromAccount)
		require.NotEmpty(t, result.result.ToAccount)
		require.NotEmpty(t, result.result.FromEntry)
		require.NotEmpty(t, result.result.ToEntry)
		require.NotEmpty(t, result.result.FromTransfer)
		require.NotEmpty(t, result.result.ToTransfer)
		// ensure records were created in the right order with the right data
		// transfers
		require.Equal(t, result.result.ToTransfer.Amount, amount)
		require.Equal(t, result.result.ToTransfer.ToAccountID, result.recipient.ID)
		require.Equal(t, result.result.ToTransfer.FromAccountID, result.sender.ID)
		require.Equal(t, result.result.FromTransfer.Amount, -amount)
		require.Equal(t, result.result.FromTransfer.FromAccountID, result.sender.ID)
		require.Equal(t, result.result.FromTransfer.ToAccountID, result.recipient.ID)
		// entries
		require.Equal(t, result.result.ToEntry.Amount, amount)
		require.Equal(t, result.result.ToEntry.AccountID, result.recipient.ID)
		require.Equal(t, result.result.FromEntry.Amount, -amount)
		require.Equal(t, result.result.FromEntry.AccountID, result.sender.ID)
		// accounts
		require.Equal(t, result.result.ToAccount.ID, result.recipient.ID)
		require.Equal(t, result.result.FromAccount.ID, result.sender.ID)
	}

	updatedSender, err := testTx.Queries.GetAccount(context.Background(), sender.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSender)
	require.Equal(t, sender.Balance, updatedSender.Balance+amount*int64(n))
	updatedRecipient, err := testTx.Queries.GetAccount(context.Background(), recipient.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRecipient)
	require.Equal(t, recipient.Balance, updatedRecipient.Balance-amount*int64(n))
}
