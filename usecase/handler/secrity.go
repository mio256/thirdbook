package handler

import (
	"context"

	"github.com/mio256/thirdbook/ui/api"
)

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	panic("implement me")
}
