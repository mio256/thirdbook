package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(ctx context.Context) *pgxpool.Pool {
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

	return dbConn
}
