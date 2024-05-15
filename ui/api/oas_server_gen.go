// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// BookingsBookingIdDelete implements DELETE /bookings/{bookingId} operation.
	//
	// Cancel an existing booking.
	//
	// DELETE /bookings/{bookingId}
	BookingsBookingIdDelete(ctx context.Context, params BookingsBookingIdDeleteParams) (BookingsBookingIdDeleteRes, error)
	// BookingsBookingIdGet implements GET /bookings/{bookingId} operation.
	//
	// Retrieve details of a specific booking by ID.
	//
	// GET /bookings/{bookingId}
	BookingsBookingIdGet(ctx context.Context, params BookingsBookingIdGetParams) (BookingsBookingIdGetRes, error)
	// BookingsBookingIdPut implements PUT /bookings/{bookingId} operation.
	//
	// Update details of an existing booking.
	//
	// PUT /bookings/{bookingId}
	BookingsBookingIdPut(ctx context.Context, req *UpdateBooking, params BookingsBookingIdPutParams) (BookingsBookingIdPutRes, error)
	// BookingsGet implements GET /bookings operation.
	//
	// Retrieve a list of all bookings.
	//
	// GET /bookings
	BookingsGet(ctx context.Context) ([]Booking, error)
	// BookingsPost implements POST /bookings operation.
	//
	// Create a new booking for a live event.
	//
	// POST /bookings
	BookingsPost(ctx context.Context, req *NewBooking) (*Booking, error)
	// UsersLoginPost implements POST /users/login operation.
	//
	// Authenticate a user and generate a token.
	//
	// POST /users/login
	UsersLoginPost(ctx context.Context, req *LoginUser) (UsersLoginPostRes, error)
	// UsersPost implements POST /users operation.
	//
	// Register a new user with the system.
	//
	// POST /users
	UsersPost(ctx context.Context, req *NewUser) (*User, error)
	// UsersUserIdDelete implements DELETE /users/{userId} operation.
	//
	// Delete an existing user.
	//
	// DELETE /users/{userId}
	UsersUserIdDelete(ctx context.Context, params UsersUserIdDeleteParams) (UsersUserIdDeleteRes, error)
	// UsersUserIdGet implements GET /users/{userId} operation.
	//
	// Retrieve details of a specific user by ID.
	//
	// GET /users/{userId}
	UsersUserIdGet(ctx context.Context, params UsersUserIdGetParams) (UsersUserIdGetRes, error)
	// UsersUserIdPut implements PUT /users/{userId} operation.
	//
	// Update details of an existing user.
	//
	// PUT /users/{userId}
	UsersUserIdPut(ctx context.Context, req *UpdateUser, params UsersUserIdPutParams) (UsersUserIdPutRes, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
