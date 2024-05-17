package handler_test

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mio256/thirdbook/pkg/testutil"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/mio256/thirdbook/usecase/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_UsersPost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		email    string
		password string
		wantErr  bool
	}{
		{
			name:     faker.Username(),
			email:    faker.Email(),
			password: faker.PASSWORD,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			dbConn := testutil.ConnectDB(t, ctx)
			h := handler.NewHandler(dbConn)
			req := &api.NewUser{
				Name:     api.OptString{Value: tt.name, Set: true},
				Email:    api.OptString{Value: tt.email, Set: true},
				Password: api.OptString{Value: tt.password, Set: true},
			}
			res, err := h.UsersPost(ctx, req)
			if !tt.wantErr {
				require.NoError(t, err)
			}
			require.Equal(t, tt.name, res.Name.Value)
			require.Equal(t, tt.email, res.Email.Value)
			require.Equal(t, tt.password, res.Password.Value)
		})
	}
}
