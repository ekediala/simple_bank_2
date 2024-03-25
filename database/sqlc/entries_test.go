package database

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createTestEntry(t *testing.T, data CreateEntryParams) Entry {
	entry, err := testQueries.CreateEntry(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func createRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	data := CreateEntryParams{
		AccountID: account.ID,
		Amount:    5000,
	}
	entry, err := testQueries.CreateEntry(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)

	data := CreateEntryParams{
		AccountID: account.ID,
		Amount:    1000,
	}

	entry := createTestEntry(t, data)

	require.Equal(t, entry.AccountID, data.AccountID)
	require.Equal(t, entry.Amount, data.Amount)
	require.NotEmpty(t, entry.CreatedAt.Time)
	require.NotEmpty(t, entry.UpdatedAt.Time)
}

func TestGetEntry(t *testing.T) {
	created := createRandomEntry(t)
	fetched, err := testQueries.GetEntry(context.Background(), created.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetched)
	require.Equal(t, fetched.ID, created.ID)
}

func TestGetEntryForUpdate(t *testing.T) {
	created := createRandomEntry(t)
	fetched, err := testQueries.GetEntryForUpdate(context.Background(), created.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetched)
	require.Equal(t, fetched.ID, created.ID)
}

func TestGetEntriesForAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	data := CreateEntryParams{
		AccountID: account.ID,
		Amount:    1000,
	}

	var numOfEntriesToCreate int = 10 // must be int. leave as is.

	var half = int64(numOfEntriesToCreate / 2)

	var wg sync.WaitGroup

	wg.Add(numOfEntriesToCreate)

	// count 1 to numOfEntriesToCreate. No need starting from zero. Just confusing for no reason imo
	for i := 1; i <= numOfEntriesToCreate; i++ {
		go func() {
			defer wg.Done()
			createTestEntry(t, data)
		}()
	}

	wg.Wait()

	// we offset entries by half the created items and expect to still have a slice of length
	// half the numOfEntriesToCreate
	results, err := testQueries.GetEntriesForAccount(context.Background(), GetEntriesForAccountParams{
		AccountID: account.ID,
		Limit:     half,
		Offset:    half,
	})

	require.NoError(t, err)
	require.NotEmpty(t, results)
	require.Len(t, results, int(half))
}

func TestUpdateEntry(t *testing.T) {
	created := createRandomEntry(t)

	entry, err := testQueries.GetEntryForUpdate(context.Background(), created.ID)
	require.NoError(t, err)

	_, err = testQueries.UpdateEntry(context.Background(), UpdateEntryParams{
		Amount: 2000,
		ID:     entry.ID,
	})
	require.NoError(t, err)

	updated, err := testQueries.GetEntry(context.Background(), created.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updated)
	require.Equal(t, updated.ID, created.ID)
	require.Equal(t, updated.Amount, int64(2000))
	require.WithinDuration(t, updated.UpdatedAt.Time, time.Now(), time.Minute)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestDeleteEntryForAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	data := CreateEntryParams{
		AccountID: account.ID,
		Amount:    1000,
	}

	numOfEntriesToCreate := 10

	var wg sync.WaitGroup

	wg.Add(numOfEntriesToCreate)

	// count 1 to numOfEntriesToCreate. No need starting from zero. Just confusing for no reason imo
	for i := 1; i <= numOfEntriesToCreate; i++ {
		go func() {
			defer wg.Done()
			createTestEntry(t, data)
		}()
	}

	wg.Wait()

	err := testQueries.DeleteEntriesForAccount(context.Background(), account.ID)
	require.NoError(t, err)

	results, err := testQueries.GetEntriesForAccount(context.Background(), GetEntriesForAccountParams{
		AccountID: account.ID,
		Limit:     1,
		Offset:    0,
	})

	require.NoError(t, err)
	require.Empty(t, results)
}
