package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEnty(t *testing.T) Entry {
	acc1 := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: acc1.ID,
		Amount: int64(float32(acc1.Balance ) * 0.1),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, acc1.ID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEnty(t)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEnty(t);

	arg := UpdateEntryParams{
		ID: entry1.ID,
		Amount: 10,
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.ID, entry1.ID)
	require.Equal(t, entry2.Amount, arg.Amount)
	require.NotZero(t, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEnty(t);

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomEnty(t)
	}

	arg := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t,entry)
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.Amount)
		require.NotZero(t, entry.CreatedAt)
	}
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEnty(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
