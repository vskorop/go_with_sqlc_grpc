package sqlc

import (
	"context"
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
