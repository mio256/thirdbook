package handler

import (
	"context"

	"github.com/mio256/thirdbook/ui/api"
)

type Handler struct{}

type SecurityHandler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func NewSecurityHandler() *SecurityHandler {
	return &SecurityHandler{}
}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	panic("not implemented")
}
