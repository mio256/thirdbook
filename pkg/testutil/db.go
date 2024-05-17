package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(t *testing.T, ctx context.Context) *pgxpool.Pool {
	t.Helper()

	var (
		dbUser         = os.Getenv("DB_USER")
		dbPwd          = os.Getenv("DB_PASS")
		unixSocketPath = os.Getenv("INSTANCE_UNIX_SOCKET")
		dbName         = os.Getenv("DB_NAME")
	)

	dbConn, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPwd, dbName, unixSocketPath))
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		dbConn.Close()
	})

	return dbConn
}
