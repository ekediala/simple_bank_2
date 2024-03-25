package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	amount := 1000

	sender := CreateRandomAccount(t)
	recipient := CreateRandomAccount(t)

	from_entry := createTestEntry(t, CreateEntryParams{
		AccountID: sender.ID,
		Amount:    int64(amount),
	})

	to_entry := createTestEntry(t, CreateEntryParams{
		AccountID: recipient.ID,
		Amount:    int64(amount),
	})

	transfer := createTestTransfer(t, CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   recipient.ID,
		FromEntryID:   from_entry.ID,
		ToEntryID:     to_entry.ID,
		Amount:        int64(amount),
	})

	return transfer
}

func createTestTransfer(t *testing.T, data CreateTransferParams) Transfer {
	transfer, err := testQueries.CreateTransfer(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	amount := 1000

	sender := CreateRandomAccount(t)
	recipient := CreateRandomAccount(t)

	from_entry := createTestEntry(t, CreateEntryParams{
		AccountID: sender.ID,
		Amount:    int64(amount),
	})

	to_entry := createTestEntry(t, CreateEntryParams{
		AccountID: recipient.ID,
		Amount:    int64(amount),
	})

	data := CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   recipient.ID,
		FromEntryID:   from_entry.ID,
		ToEntryID:     to_entry.ID,
		Amount:        int64(amount),
	}

	transfer := createTestTransfer(t, data)

	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.Amount, data.Amount)
	require.Equal(t, transfer.FromAccountID, data.FromAccountID)
	require.Equal(t, transfer.ToAccountID, data.ToAccountID)
	require.Equal(t, transfer.FromEntryID, data.FromEntryID)
	require.Equal(t, transfer.ToEntryID, data.ToEntryID)
	require.NotEmpty(t, transfer.CreatedAt)
	require.NotEmpty(t, transfer.UpdatedAt)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	result, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, transfer.ID, result.ID)
}

func TestGetTransferForUpdate(t *testing.T) {
	transfer := createRandomTransfer(t)
	result, err := testQueries.GetTransferForUpdate(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, transfer.ID, result.ID)
}

func TestUpdateTransferAmount(t *testing.T) {
	transfer := createRandomTransfer(t)

	transfer, err := testQueries.GetTransferForUpdate(context.Background(), transfer.ID)
	require.NoError(t, err)

	_, err = testQueries.UpdateTransferAmount(context.Background(), UpdateTransferAmountParams{
		TransferAmount: transfer.Amount * 2,
		TransferID:     transfer.ID,
	})
	require.NoError(t, err)

	updated, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NotEmpty(t, updated)
	require.NoError(t, err)

	require.Equal(t, transfer.ID, updated.ID)
	require.Equal(t, transfer.Amount*2, updated.Amount)
	require.WithinDuration(t, updated.UpdatedAt.Time, time.Now(), time.Minute)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	result, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.Empty(t, result)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
