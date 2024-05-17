package handler

import (
	"context"
	"net/mail"
	"strconv"

	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/taxio/errors"
)

func (h *Handler) UsersLoginPost(ctx context.Context, req *api.LoginUser) (api.UsersLoginPostRes, error) {
	panic("not implemented")
}

func (h *Handler) UsersPost(ctx context.Context, req *api.NewUser) (*api.User, error) {
	email := req.Email.Value
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	user, err := h.repo.CreateUser(ctx, rdb.CreateUserParams{
		Name:     req.Name.Value,
		Email:    email,
		Password: req.Password.Value, // TODO: hash
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.User{
		ID:        api.OptString{Value: strconv.FormatInt(user.ID, 10), Set: true},
		Name:      api.OptString{Value: user.Name, Set: true},
		Email:     api.OptString{Value: user.Email, Set: true},
		Password:  api.OptString{Value: user.Password, Set: true},
		CreatedAt: api.OptDateTime{Value: user.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: user.UpdatedAt.Time, Set: true},
	}, nil
}

func (h *Handler) UsersUserIdDelete(ctx context.Context, params api.UsersUserIdDeleteParams) (api.UsersUserIdDeleteRes, error) {
	panic("not implemented")
}

func (h *Handler) UsersUserIdGet(ctx context.Context, params api.UsersUserIdGetParams) (api.UsersUserIdGetRes, error) {
	panic("not implemented")
}

func (h *Handler) UsersUserIdPut(ctx context.Context, req *api.UpdateUser, params api.UsersUserIdPutParams) (api.UsersUserIdPutRes, error) {
	panic("not implemented")
}
