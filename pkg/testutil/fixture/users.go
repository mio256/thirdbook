package fixture

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/stretchr/testify/require"
)

func CreateUser(t *testing.T, ctx context.Context, db rdb.DBTX, f func(v *rdb.User)) *rdb.User {
	t.Helper()

	target := &rdb.User{
		Name:     faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	if f != nil {
		f(target)
	}

	created, err := rdb.New(db).TestCreateUser(ctx, rdb.TestCreateUserParams{
		Name:     target.Name,
		Email:    target.Email,
		Password: target.Password,
	})

	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, rdb.New(db).TestDeleteUser(ctx, created.ID))
	})

	return &created
}
