package infra_test

import (
	"context"
	"testing"

	"github.com/mio256/thirdbook/pkg/infra"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
)

func Test_ConnectDB(t *testing.T) {
	ctx := context.Background()
	dbConn := infra.ConnectDB(ctx)
	repo := rdb.New(dbConn)
	if err := repo.Ping(ctx); err != nil {
		panic(err)
	}
	defer dbConn.Close()
}
