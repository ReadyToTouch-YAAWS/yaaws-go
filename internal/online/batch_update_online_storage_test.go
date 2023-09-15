package main

import (
	"context"
	"testing"

	postgresql "github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/stretchr/testify/require"
)

func TestBatchUpdateOnlineStorage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	storage := NewBatchUpdateOnlineStorage(postgresql.NewSqlcRepository(connection))

	testOnlineStorage(t, storage)
}

func BenchmarkBatchUpdateOnlineStorage(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	ctx := context.Background()

	connection, err := pgxpool.New(ctx, dataSourceName)
	require.NoError(b, err)
	defer connection.Close()

	storage := NewBatchUpdateOnlineStorage(postgresql.NewSqlcRepository(connection))

	benchmarkOnlineStorage(b, storage)
}
