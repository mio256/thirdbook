package main

import (
	"log"
	"net/http"

	"github.com/mio256/thirdbook/ui/api"
	"github.com/mio256/thirdbook/usecase/handler"
)

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

func main() {
	h := handler.NewHandler()
	s := handler.NewSecurityHandler()
	srv, err := api.NewServer(h, s)
	if err != nil {
		log.Fatal(err)
	}

	// Wrap the server with the CORS middleware
	corsHandler := enableCORS(srv)

	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal(err)
	}
}
