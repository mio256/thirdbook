// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// BookingsBookingIDGet implements GET /bookings/{bookingID} operation.
//
// Retrieve details of a specific booking by ID.
//
// GET /bookings/{bookingID}
func (UnimplementedHandler) BookingsBookingIDGet(ctx context.Context, params BookingsBookingIDGetParams) (r BookingsBookingIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// BookingsBookingIDPut implements PUT /bookings/{bookingID} operation.
//
// Cancel an existing booking.
//
// PUT /bookings/{bookingID}
func (UnimplementedHandler) BookingsBookingIDPut(ctx context.Context, params BookingsBookingIDPutParams) (r BookingsBookingIDPutRes, _ error) {
	return r, ht.ErrNotImplemented
}

// BookingsGet implements GET /bookings operation.
//
// Retrieve a list of all bookings that meets the conditions.
//
// GET /bookings
func (UnimplementedHandler) BookingsGet(ctx context.Context, params BookingsGetParams) (r []Booking, _ error) {
	return r, ht.ErrNotImplemented
}

// BookingsPost implements POST /bookings operation.
//
// Create a new booking for a live event.
//
// POST /bookings
func (UnimplementedHandler) BookingsPost(ctx context.Context, req *NewBooking) (r *Booking, _ error) {
	return r, ht.ErrNotImplemented
}

// PingGet implements GET /ping operation.
//
// Check if the server is running.
//
// GET /ping
func (UnimplementedHandler) PingGet(ctx context.Context) (r *PingGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersLoginPost implements POST /users/login operation.
//
// Authenticate a user and generate a token.
//
// POST /users/login
func (UnimplementedHandler) UsersLoginPost(ctx context.Context, req *LoginUser) (r UsersLoginPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersPost implements POST /users operation.
//
// Register a new user with the system.
//
// POST /users
func (UnimplementedHandler) UsersPost(ctx context.Context, req *NewUser) (r *User, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersUserIDDelete implements DELETE /users/{userID} operation.
//
// Delete an existing user.
//
// DELETE /users/{userID}
func (UnimplementedHandler) UsersUserIDDelete(ctx context.Context, params UsersUserIDDeleteParams) (r UsersUserIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UsersUserIDGet implements GET /users/{userID} operation.
//
// Retrieve details of a specific user by ID.
//
// GET /users/{userID}
func (UnimplementedHandler) UsersUserIDGet(ctx context.Context, params UsersUserIDGetParams) (r UsersUserIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
