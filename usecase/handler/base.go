package handler

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/ui/api"
)

type Handler struct {
	dbConn *pgxpool.Pool
	repo   *rdb.Queries
}

type SecurityHandler struct{}

func NewHandler(dbConn *pgxpool.Pool) *Handler {
	return &Handler{
		dbConn: dbConn,
		repo:   rdb.New(dbConn),
	}
}

func NewSecurityHandler() *SecurityHandler {
	return &SecurityHandler{}
}

func (h *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	panic("not implemented")
}
