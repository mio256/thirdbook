package handler

import (
	"context"

	"github.com/mio256/thirdbook/ui/api"
)

func (h *Handler) PingGet(ctx context.Context) (*api.PingGetOK, error) {
	res := &api.PingGetOK{Message: api.OptString{Value: "pong", Set: true}}
	return res, nil
}
