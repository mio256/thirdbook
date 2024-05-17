package handler

import (
	"context"
	"net/mail"
	"strconv"
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
		ID:        api.OptString{Value: strconv.FormatInt(user.ID, 10), Set: true},
		Name:      api.OptString{Value: user.Name, Set: true},
		Email:     api.OptString{Value: user.Email, Set: true},
		Password:  api.OptString{Value: user.Password, Set: true},
		CreatedAt: api.OptDateTime{Value: user.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: user.UpdatedAt.Time, Set: true},
	}, nil
}

func (h *Handler) UsersUserIdDelete(ctx context.Context, params api.UsersUserIdDeleteParams) (api.UsersUserIdDeleteRes, error) {
	id, err := strconv.ParseUint(params.UserId, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	_, err = h.repo.GetUser(ctx, int64(id))
	if err != nil {
		return &api.UsersUserIdDeleteNotFound{}, errors.Wrap(err)
	}

	err = h.repo.DeleteUser(ctx, int64(id))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.UsersUserIdDeleteNoContent{}, nil
}

func (h *Handler) UsersUserIdGet(ctx context.Context, params api.UsersUserIdGetParams) (api.UsersUserIdGetRes, error) {
	id, err := strconv.ParseUint(params.UserId, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	user, err := h.repo.GetUser(ctx, int64(id))
	if err != nil {
		return &api.UsersUserIdGetNotFound{}, errors.Wrap(err)
	}

	return &api.User{
		ID:        api.OptString{Value: strconv.FormatInt(user.ID, 10), Set: true},
		Name:      api.OptString{Value: user.Name, Set: true},
		Email:     api.OptString{Value: user.Email, Set: true},
		CreatedAt: api.OptDateTime{Value: user.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: user.UpdatedAt.Time, Set: true},
	}, nil
}
