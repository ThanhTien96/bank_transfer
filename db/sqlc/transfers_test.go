package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAcc := createRandomAccount(t)
	toAcc := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAcc.ID,
		ToAccountID:   toAcc.ID,
		Amount:        int64(float32(fromAcc.Balance) * 0.1),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, fromAcc.ID)
	require.Equal(t, transfer.ToAccountID, toAcc.ID)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID: transfer.ID,
		Amount: 20,
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer.ID)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer2.Amount, arg.Amount)

	require.WithinDuration(t, transfer2.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer.ID)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer2.Amount, transfer.Amount)
	require.WithinDuration(t, transfer2.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err  := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t,transfer2)
}