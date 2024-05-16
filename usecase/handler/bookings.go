package handler

import (
	"context"

	"github.com/mio256/thirdbook/ui/api"
)

func (h *Handler) BookingsBookingIdDelete(ctx context.Context, params api.BookingsBookingIdDeleteParams) (api.BookingsBookingIdDeleteRes, error) {
	panic("not implemented")
}

func (h *Handler) BookingsBookingIdGet(ctx context.Context, params api.BookingsBookingIdGetParams) (api.BookingsBookingIdGetRes, error) {
	panic("not implemented")
}

func (h *Handler) BookingsBookingIdPut(ctx context.Context, req *api.UpdateBooking, params api.BookingsBookingIdPutParams) (api.BookingsBookingIdPutRes, error) {
	panic("not implemented")
}

func (h *Handler) BookingsGet(ctx context.Context) ([]api.Booking, error) {
	panic("not implemented")
}

func (h *Handler) BookingsPost(ctx context.Context, req *api.NewBooking) (*api.Booking, error) {
	panic("not implemented")
}
