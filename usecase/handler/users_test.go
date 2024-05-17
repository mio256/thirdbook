package handler_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/pkg/testutil"
	"github.com/mio256/thirdbook/pkg/testutil/fixture"
	"github.com/mio256/thirdbook/pkg/util"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/mio256/thirdbook/usecase/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_UsersPost(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	tests := []struct {
		name     string
		email    string
		password string
		wantErr  bool
	}{
		{
			name:     "valid",
			email:    faker.Email(),
			password: faker.Password(),
			wantErr:  false,
		},
		{
			name:     "invalid-email",
			email:    "invalid-email",
			password: faker.Password(),
			wantErr:  true,
		},
		{
			name:     "empty-password",
			email:    faker.Email(),
			password: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)
			req := &api.NewUser{
				Name:     api.OptString{Value: tt.name, Set: true},
				Email:    api.OptString{Value: tt.email, Set: true},
				Password: api.OptString{Value: tt.password, Set: true},
			}
			res, err := h.UsersPost(ctx, req)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.name, res.Name.Value)
				require.Equal(t, tt.email, res.Email.Value)
				require.NotEqual(t, tt.password, res.Password.Value)
				require.NoError(t, util.CompareHashAndPassword(res.Password.Value, tt.password))
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestHandler_UsersUserIdDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	tests := []struct {
		name    string
		user    *rdb.User
		wantErr bool
	}{
		{
			name:    "valid",
			user:    fixture.CreateUser(t, ctx, dbConn, nil),
			wantErr: false,
		},
		{
			name:    "invalid-user",
			user:    &rdb.User{ID: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)

			params := api.UsersUserIdDeleteParams{UserId: strconv.FormatInt(tt.user.ID, 10)}
			res, err := h.UsersUserIdDelete(ctx, params)
			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, &api.UsersUserIdDeleteNoContent{}, res)
			} else {
				require.Error(t, err)
				require.Equal(t, &api.UsersUserIdDeleteNotFound{}, res)
			}
		})
	}
}

func TestHandler_UsersUserIdGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	tests := []struct {
		name    string
		user    *rdb.User
		wantErr bool
	}{
		{
			name:    "valid",
			user:    fixture.CreateUser(t, ctx, dbConn, nil),
			wantErr: false,
		},
		{
			name:    "invalid-user",
			user:    &rdb.User{ID: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)

			params := api.UsersUserIdGetParams{UserId: strconv.FormatInt(tt.user.ID, 10)}
			res, err := h.UsersUserIdGet(ctx, params)
			if !tt.wantErr {
				require.NoError(t, err)
				res200, err := res.(*api.User)
				require.True(t, err)
				require.Equal(t, strconv.FormatInt(tt.user.ID, 10), res200.ID.Value)
				require.Equal(t, tt.user.Name, res200.Name.Value)
				require.Equal(t, tt.user.Email, res200.Email.Value)
				require.Equal(t, tt.user.Password, res200.Password.Value)
			} else {
				require.Error(t, err)
				require.Equal(t, &api.UsersUserIdGetNotFound{}, res)
			}
		})
	}
}
