package handler

import (
	"context"
	"net/mail"
	"unicode/utf8"

	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/pkg/util"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/taxio/errors"
)

func (h *Handler) UsersLoginPost(ctx context.Context, req *api.LoginUser) (api.UsersLoginPostRes, error) {
	user, err := h.repo.GetUserByEmail(ctx, req.Email.Value)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	err = util.CompareHashAndPassword(user.Password, req.Password.Value)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	token, err := util.GenerateToken("thirdbook", util.UserToken{
		ID:    uint64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.AuthToken{
		Token: api.OptString{Value: token, Set: true},
	}, nil
}

func (h *Handler) UsersPost(ctx context.Context, req *api.NewUser) (*api.User, error) {
	email := req.Email.Value
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if utf8.RuneCountInString(req.Password.Value) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	password, err := util.GeneratePasswordHash(req.Password.Value)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	user, err := h.repo.CreateUser(ctx, rdb.CreateUserParams{
		Name:     req.Name.Value,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.User{
		ID:        api.OptInt64{Value: user.ID, Set: true},
		Name:      api.OptString{Value: user.Name, Set: true},
		Email:     api.OptString{Value: user.Email, Set: true},
		Password:  api.OptString{Value: user.Password, Set: true},
		CreatedAt: api.OptDateTime{Value: user.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: user.UpdatedAt.Time, Set: true},
	}, nil
}

func (h *Handler) UsersUserIDDelete(ctx context.Context, params api.UsersUserIDDeleteParams) (api.UsersUserIDDeleteRes, error) {
	_, err := h.repo.GetUser(ctx, params.UserID)
	if err != nil {
		return &api.UsersUserIDDeleteNotFound{}, errors.Wrap(err)
	}

	err = h.repo.DeleteUser(ctx, params.UserID)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.UsersUserIDDeleteNoContent{}, nil
}

func (h *Handler) UsersUserIDGet(ctx context.Context, params api.UsersUserIDGetParams) (api.UsersUserIDGetRes, error) {
	user, err := h.repo.GetUser(ctx, params.UserID)
	if err != nil {
		return &api.UsersUserIDGetNotFound{}, errors.Wrap(err)
	}

	return &api.User{
		ID:        api.OptInt64{Value: user.ID, Set: true},
		Name:      api.OptString{Value: user.Name, Set: true},
		Email:     api.OptString{Value: user.Email, Set: true},
		CreatedAt: api.OptDateTime{Value: user.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: user.UpdatedAt.Time, Set: true},
	}, nil
}
