package fixture

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mio256/thirdbook/pkg/infra/rdb"
	"github.com/stretchr/testify/require"
)

func CreateBooking(t *testing.T, ctx context.Context, db rdb.DBTX, f func(v *rdb.Booking)) *rdb.Booking {
	t.Helper()

	user := CreateUser(t, ctx, db, nil)

	target := &rdb.Booking{
		Name: faker.Username(),
		StartTime: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		EndTime: pgtype.Timestamp{
			Time:  time.Now().Add(time.Hour),
			Valid: true,
		},
		UserID: user.ID,
		Status: rdb.BookingTypePending,
	}

	if f != nil {
		f(target)
	}

	created, err := rdb.New(db).TestCreateBooking(ctx, rdb.TestCreateBookingParams{
		Name:      target.Name,
		StartTime: target.StartTime,
		EndTime:   target.EndTime,
		UserID:    target.UserID,
		Status:    target.Status,
	})

	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, rdb.New(db).TestDeleteBooking(ctx, created.ID))
	})

	return &created
}
