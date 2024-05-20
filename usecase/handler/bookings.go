package handler

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/taxio/errors"
)

func (h *Handler) BookingsBookingIDPut(ctx context.Context, params api.BookingsBookingIDPutParams) (api.BookingsBookingIDPutRes, error) {
	s, err := h.repo.GetBookingStatus(ctx, params.BookingID)
	if err != nil {
		return &api.BookingsBookingIDPutNotFound{}, errors.Wrap(err)
	}
	if s != rdb.BookingTypePending {
		return nil, errors.New(fmt.Sprintf("this booking is not pending: %s", string(s)))
	}
	b, err := h.repo.UpdateBookingStatus(ctx, rdb.UpdateBookingStatusParams{
		Status: rdb.BookingTypeCanceled,
		ID:     params.BookingID,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &api.Booking{
		ID:    api.OptInt64{Value: b.ID, Set: true},
		Name:  api.OptString{Value: b.Name, Set: true},
		Start: api.OptDateTime{Value: b.StartTime.Time, Set: true},
		End:   api.OptDateTime{Value: b.EndTime.Time, Set: true},
		User:  api.OptInt64{Value: b.UserID, Set: true},
		Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
			Value: string(b.Status),
			Set:   true,
		}}, Set: true},
		CreatedAt: api.OptDateTime{Value: b.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: b.UpdatedAt.Time, Set: true},
	}, nil
}

func (h *Handler) BookingsBookingIDGet(ctx context.Context, params api.BookingsBookingIDGetParams) (api.BookingsBookingIDGetRes, error) {
	b, err := h.repo.GetBooking(ctx, params.BookingID)
	if err != nil {
		return &api.BookingsBookingIDGetNotFound{}, errors.Wrap(err)
	}
	return &api.Booking{
		ID:    api.OptInt64{Value: b.ID, Set: true},
		Name:  api.OptString{Value: b.Name, Set: true},
		Start: api.OptDateTime{Value: b.StartTime.Time, Set: true},
		End:   api.OptDateTime{Value: b.EndTime.Time, Set: true},
		User:  api.OptInt64{Value: b.UserID, Set: true},
		Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
			Value: string(b.Status),
			Set:   true,
		}}, Set: true},
		CreatedAt: api.OptDateTime{Value: b.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: b.UpdatedAt.Time, Set: true},
	}, nil
}

func (h *Handler) BookingsGet(ctx context.Context, params api.BookingsGetParams) ([]api.Booking, error) {
	bs, err := h.repo.GetBookings(ctx, rdb.GetBookingsParams{
		Status:     rdb.BookingType(params.Status.Value.Status.Value),
		UserID:     params.User.Value,
		StartLimit: pgtype.Timestamp{Time: params.Start.Value, Valid: true},
		EndLimit:   pgtype.Timestamp{Time: params.End.Value, Valid: true},
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}
	var res []api.Booking
	for _, b := range bs {
		res = append(res, api.Booking{
			ID:    api.OptInt64{Value: b.ID, Set: true},
			Name:  api.OptString{Value: b.Name, Set: true},
			Start: api.OptDateTime{Value: b.StartTime.Time, Set: true},
			End:   api.OptDateTime{Value: b.EndTime.Time, Set: true},
			User:  api.OptInt64{Value: b.UserID, Set: true},
			Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
				Value: string(b.Status),
				Set:   true,
			}}, Set: true},
			CreatedAt: api.OptDateTime{Value: b.CreatedAt.Time, Set: true},
			UpdatedAt: api.OptDateTime{Value: b.UpdatedAt.Time, Set: true},
		})
	}
	return res, nil
}

func (h *Handler) BookingsPost(ctx context.Context, req *api.NewBooking) (*api.Booking, error) {
	b, err := h.repo.CreateBooking(ctx, rdb.CreateBookingParams{
		Name:      req.Name.Value,
		StartTime: pgtype.Timestamp{Time: req.Start.Value, Valid: true},
		EndTime:   pgtype.Timestamp{Time: req.End.Value, Valid: true},
		UserID:    req.User.Value,
		Status:    rdb.BookingTypePending,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &api.Booking{
		ID:    api.OptInt64{Value: b.ID, Set: true},
		Name:  api.OptString{Value: b.Name, Set: true},
		Start: api.OptDateTime{Value: b.StartTime.Time, Set: true},
		End:   api.OptDateTime{Value: b.EndTime.Time, Set: true},
		User:  api.OptInt64{Value: b.UserID, Set: true},
		Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
			Value: string(b.Status),
			Set:   true,
		}}, Set: true},
		CreatedAt: api.OptDateTime{Value: b.CreatedAt.Time, Set: true},
		UpdatedAt: api.OptDateTime{Value: b.UpdatedAt.Time, Set: true},
	}, nil
}
