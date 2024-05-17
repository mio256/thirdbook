package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/mio256/thirdbook/pkg/infra"
	"github.com/mio256/thirdbook/ui/api"
	"github.com/mio256/thirdbook/usecase/handler"
	"github.com/ogen-go/ogen/middleware"
)

func main() {
	var slogOpts slog.HandlerOptions
	slogOpts.Level = slog.LevelDebug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slogOpts))
	slog.SetDefault(logger)

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	dbConn := infra.ConnectDB(ctx)
	h := handler.NewHandler(dbConn)
	s := handler.NewSecurityHandler()
	srv, err := api.NewServer(h, s, api.WithMiddleware(loggingMiddleware()))
	if err != nil {
		return err
	}

	// Wrap the server with the CORS middleware
	corsHandler := enableCORS(srv)

	addr := ":8080"
	slog.InfoContext(ctx, "Starting HTTP Server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		return err
	}

	return nil
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins, adjust as necessary
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Check if the request is for CORS OPTIONS (pre-flight)
		if r.Method == "OPTIONS" {
			// Just add headers and send response
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the next handler
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		slog.DebugContext(
			req.Context, fmt.Sprintf("[http] <-- %s", req.OperationName),
			slog.String("path", req.Raw.URL.String()),
			slog.Any("body", req.Body),
		)

		resp, err := next(req)
		if err != nil {
			slog.DebugContext(
				req.Context, fmt.Sprintf("FAIL %s", req.OperationName),
				slog.String("message", err.Error()),
			)
		} else {
			slog.DebugContext(
				req.Context, fmt.Sprintf("[http] --> %s", req.OperationName),
				slog.Any("body", resp.Type),
			)
		}

		return resp, err
	}
}
