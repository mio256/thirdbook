package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mio256/thirdbook/pkg/util"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/taxio/errors"
)

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	token, err := util.ParseToken(t.Token)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	id, err := strconv.ParseUint(fmt.Sprintf("%.0f", token.Claims.(jwt.MapClaims)["user_id"].(float64)), 10, 64)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	email := token.Claims.(jwt.MapClaims)["email"].(string)

	name := token.Claims.(jwt.MapClaims)["name"].(string)

	user := util.UserToken{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return context.WithValue(ctx, "user", user), nil
}
