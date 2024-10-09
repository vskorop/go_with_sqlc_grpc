package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vskorop/go_with_sqlc_grpc/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountRandom := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), accountRandom.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.Equal(t, account.ID, accountRandom.ID)
	require.Equal(t, account.Owner, accountRandom.Owner)
	require.Equal(t, account.Balance, accountRandom.Balance)
	require.Equal(t, account.Currency, accountRandom.Currency)
	require.WithinDuration(t, account.CreatedAt.Time, accountRandom.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accountRandom := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      accountRandom.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testQueries.GetAccount(context.Background(), accountRandom.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, account.ID, accountRandom.ID)
	require.Equal(t, account.Owner, accountRandom.Owner)
	require.Equal(t, account.Currency, accountRandom.Currency)

	require.NotEqual(t, arg.Balance, accountRandom.Balance)

	require.WithinDuration(t, accountRandom.CreatedAt.Time, account.CreatedAt.Time, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	getAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getAccount)
}
func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
func TestGetAccountNotFound(t *testing.T) {
	account, err := testQueries.GetAccount(context.Background(), -1)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}
