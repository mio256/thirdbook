package handler_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/mio256/thirdbook/pkg/testutil"
	"github.com/mio256/thirdbook/pkg/testutil/fixture"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/mio256/thirdbook/usecase/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_BookingsBookingIDPut(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	tests := []struct {
		name    string
		booking *rdb.Booking
		wantErr bool
	}{
		{
			name:    "valid",
			booking: fixture.CreateBooking(t, ctx, dbConn, nil),
			wantErr: false,
		},
		{
			name:    "not-found",
			booking: &rdb.Booking{ID: 0},
			wantErr: true,
		},
		{
			name: "invalid-status",
			booking: fixture.CreateBooking(t, ctx, dbConn, func(v *rdb.Booking) {
				v.Status = rdb.BookingTypeApproved
			}),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)

			res, err := h.BookingsBookingIDPut(ctx, api.BookingsBookingIDPutParams{
				BookingID: tt.booking.ID,
			})

			if !tt.wantErr {
				require.NoError(t, err)
				res200, err := res.(*api.Booking)
				require.True(t, err)
				require.Equal(t, tt.booking.ID, res200.ID.Value)
				require.Equal(t, tt.booking.Name, res200.Name.Value)
				require.Equal(t, tt.booking.StartTime.Time, res200.Start.Value)
				require.Equal(t, tt.booking.EndTime.Time, res200.End.Value)
				require.Equal(t, tt.booking.UserID, res200.User.Value)
				require.Equal(t, string(rdb.BookingTypeCanceled), res200.Status.Value.Status.Value)
				require.WithinDuration(t, time.Now(), res200.CreatedAt.Value, time.Minute)
				require.WithinDuration(t, time.Now(), res200.UpdatedAt.Value, time.Minute)
			} else {
				require.Error(t, err)
				if res != nil {
					require.Equal(t, &api.BookingsBookingIDPutNotFound{}, res)
				}
			}
		})
	}
}

func TestHandler_BookingsBookingIDGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	tests := []struct {
		name    string
		booking *rdb.Booking
		wantErr bool
	}{
		{
			name:    "valid",
			booking: fixture.CreateBooking(t, ctx, dbConn, nil),
			wantErr: false,
		},
		{
			name:    "not-found",
			booking: &rdb.Booking{ID: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)

			res, err := h.BookingsBookingIDGet(ctx, api.BookingsBookingIDGetParams{
				BookingID: tt.booking.ID,
			})

			if !tt.wantErr {
				require.NoError(t, err)
				res200, ok := res.(*api.Booking)
				require.True(t, ok)
				require.Equal(t, tt.booking.ID, res200.ID.Value)
				require.Equal(t, tt.booking.Name, res200.Name.Value)
				require.Equal(t, tt.booking.StartTime.Time, res200.Start.Value)
				require.Equal(t, tt.booking.EndTime.Time, res200.End.Value)
				require.Equal(t, tt.booking.UserID, res200.User.Value)
				require.WithinDuration(t, time.Now(), res200.CreatedAt.Value, time.Minute)
				require.WithinDuration(t, time.Now(), res200.UpdatedAt.Value, time.Minute)
				require.Equal(t, string(tt.booking.Status), res200.Status.Value.Status.Value)
			} else {
				require.Error(t, err)
				require.Equal(t, &api.BookingsBookingIDGetNotFound{}, res)
			}
		})
	}
}

func TestHandler_BookingsGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbConn := testutil.ConnectDB(t, ctx)

	user := fixture.CreateUser(t, ctx, dbConn, nil)
	booking := fixture.CreateBooking(t, ctx, dbConn, func(b *rdb.Booking) {
		b.Name = fmt.Sprintf("booking-%d", user.ID)
		b.Status = rdb.BookingTypePending
		b.UserID = user.ID
	})

	tests := []struct {
		name    string
		params  api.BookingsGetParams
		user    *rdb.User
		want    []api.Booking
		wantErr bool
	}{
		{
			name: "valid",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
			},
			want: []api.Booking{
				{
					ID:    api.OptInt64{Value: booking.ID, Set: true},
					Name:  api.OptString{Value: fmt.Sprintf("booking-%d", user.ID), Set: true},
					Start: api.OptDateTime{Value: booking.StartTime.Time, Set: true},
					End:   api.OptDateTime{Value: booking.EndTime.Time, Set: true},
					User:  api.OptInt64{Value: user.ID, Set: true},
					Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
						Value: string(rdb.BookingTypePending),
						Set:   true,
					}}, Set: true},
					CreatedAt: api.OptDateTime{Value: booking.CreatedAt.Time, Set: true},
					UpdatedAt: api.OptDateTime{Value: booking.UpdatedAt.Time, Set: true},
				},
			},
			wantErr: false,
		},
		{
			name: "valid-start",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
				Start:  api.OptDateTime{Value: booking.StartTime.Time.Add(time.Hour / 2), Set: true},
			},
			want: []api.Booking{
				{
					ID:    api.OptInt64{Value: booking.ID, Set: true},
					Name:  api.OptString{Value: fmt.Sprintf("booking-%d", user.ID), Set: true},
					Start: api.OptDateTime{Value: booking.StartTime.Time, Set: true},
					End:   api.OptDateTime{Value: booking.EndTime.Time, Set: true},
					User:  api.OptInt64{Value: user.ID, Set: true},
					Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
						Value: string(rdb.BookingTypePending),
						Set:   true,
					}}, Set: true},
					CreatedAt: api.OptDateTime{Value: booking.CreatedAt.Time, Set: true},
					UpdatedAt: api.OptDateTime{Value: booking.UpdatedAt.Time, Set: true},
				},
			},
			wantErr: false,
		},
		{
			name: "valid-end",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
				End:    api.OptDateTime{Value: booking.StartTime.Time.Add(-time.Hour / 2), Set: true},
			},
			want: []api.Booking{
				{
					ID:    api.OptInt64{Value: booking.ID, Set: true},
					Name:  api.OptString{Value: fmt.Sprintf("booking-%d", user.ID), Set: true},
					Start: api.OptDateTime{Value: booking.StartTime.Time, Set: true},
					End:   api.OptDateTime{Value: booking.EndTime.Time, Set: true},
					User:  api.OptInt64{Value: user.ID, Set: true},
					Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
						Value: string(rdb.BookingTypePending),
						Set:   true,
					}}, Set: true},
					CreatedAt: api.OptDateTime{Value: booking.CreatedAt.Time, Set: true},
					UpdatedAt: api.OptDateTime{Value: booking.UpdatedAt.Time, Set: true},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid-start",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
				Start:  api.OptDateTime{Value: booking.StartTime.Time.Add(-time.Hour / 2), Set: true},
			},
			want: []api.Booking{
				{
					ID:    api.OptInt64{Value: booking.ID, Set: true},
					Name:  api.OptString{Value: fmt.Sprintf("booking-%d", user.ID), Set: true},
					Start: api.OptDateTime{Value: booking.StartTime.Time, Set: true},
					End:   api.OptDateTime{Value: booking.EndTime.Time, Set: true},
					User:  api.OptInt64{Value: user.ID, Set: true},
					Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
						Value: string(rdb.BookingTypePending),
						Set:   true,
					}}, Set: true},
					CreatedAt: api.OptDateTime{Value: booking.CreatedAt.Time, Set: true},
					UpdatedAt: api.OptDateTime{Value: booking.UpdatedAt.Time, Set: true},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid-end",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
				End:    api.OptDateTime{Value: booking.StartTime.Time.Add(time.Hour / 2), Set: true},
			},
			want: []api.Booking{
				{
					ID:    api.OptInt64{Value: booking.ID, Set: true},
					Name:  api.OptString{Value: fmt.Sprintf("booking-%d", user.ID), Set: true},
					Start: api.OptDateTime{Value: booking.StartTime.Time, Set: true},
					End:   api.OptDateTime{Value: booking.EndTime.Time, Set: true},
					User:  api.OptInt64{Value: user.ID, Set: true},
					Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{
						Value: string(rdb.BookingTypePending),
						Set:   true,
					}}, Set: true},
					CreatedAt: api.OptDateTime{Value: booking.CreatedAt.Time, Set: true},
					UpdatedAt: api.OptDateTime{Value: booking.UpdatedAt.Time, Set: true},
				},
			},
			wantErr: false,
		},
		{
			name: "not-found",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: string(rdb.BookingTypePending), Set: true}}, Set: true},
				User:   api.OptInt64{Value: 0, Set: true},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid-input-status",
			params: api.BookingsGetParams{
				Status: api.OptBookingStatus{Value: api.BookingStatus{Status: api.OptString{Value: "invalid-status", Set: true}}, Set: true},
				User:   api.OptInt64{Value: user.ID, Set: true},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewHandler(dbConn)

			res, err := h.BookingsGet(ctx, tt.params)

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, res)
			} else {
				require.Error(t, err)
				require.Equal(t, tt.want, res)
			}
		})
	}
}
