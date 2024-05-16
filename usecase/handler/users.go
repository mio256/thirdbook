package handler

import (
	"context"

	"github.com/mio256/thirdbook/ui/api"
)

func (h *Handler) UsersLoginPost(ctx context.Context, req *api.LoginUser) (api.UsersLoginPostRes, error) {
	panic("not implemented")
}

func (h *Handler) UsersPost(ctx context.Context, req *api.NewUser) (*api.User, error) {
	panic("not implemented")
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
