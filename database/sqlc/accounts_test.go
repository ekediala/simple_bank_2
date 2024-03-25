package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func CreateTestAccount(t *testing.T, data CreateAccountParams) Account {
	account, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    data.Owner,
		Currency: data.Currency,
	})

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func CreateRandomAccount(t *testing.T) Account {
	fake := faker.New()
	owner := fake.Person().Name()

	data := CreateAccountParams{
		Owner:    owner,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    data.Owner,
		Currency: data.Currency,
	})

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func TestCreateAccount(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	account := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	require.Equal(t, account.Balance, int64(0))
	require.Equal(t, account.Owner, owner)
	require.Equal(t, account.Currency, "NGN")
}

func TestGetAccount(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	accountCreated := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	account, err := testQueries.GetAccountForUpdate(context.Background(), accountCreated.ID)

	require.NoError(t, err)
	require.Equal(t, account.Owner, accountCreated.Owner)
	require.Equal(t, account.Currency, accountCreated.Currency)

}

func TestGetAccountForUpdate(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	accountCreated := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	account, err := testQueries.GetAccount(context.Background(), accountCreated.ID)

	require.NoError(t, err)
	require.Equal(t, account.Owner, accountCreated.Owner)
	require.Equal(t, account.Currency, accountCreated.Currency)

}

func TestDeleteAccount(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	account := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	acc, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc)
}

func TestUpdateAccountOwner(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	created := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	account, err := testQueries.GetAccountForUpdate(context.Background(), created.ID)
	require.NoError(t, err)

	newOwner := fake.Person().Name()

	updated, err := testQueries.UpdateAccountOwner(context.Background(), UpdateAccountOwnerParams{
		ID:    account.ID,
		Owner: newOwner,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updated)
	// expect owner to be new owner and updated at to be now
	require.Equal(t, updated.Owner, newOwner)
	require.WithinDuration(t, updated.UpdatedAt.Time, time.Now(), 10*time.Second)

}

func TestUpdateAccountBalance(t *testing.T) {
	fake := faker.New()
	owner := fake.Person().Name()

	account := CreateTestAccount(t, CreateAccountParams{
		Owner:    owner,
		Currency: "NGN",
	})

	newBalance := int64(2000)

	updated, err := testQueries.UpdateAccountBalance(context.Background(), UpdateAccountBalanceParams{
		ID:      account.ID,
		Amount: newBalance,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updated)
	// expect owner to be new owner and updated at to be now
	require.Equal(t, updated.Balance, newBalance)
	require.WithinDuration(t, updated.UpdatedAt.Time, time.Now(), 10*time.Second)

}
