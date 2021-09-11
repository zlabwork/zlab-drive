package middleware

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Request-Id") == "" {
			id := uuid.New()
			w.Header().Set("Request-Id", id.String())
		}

		// Do stuff here
		log.Println(r.Header.Get("Request-Id"), r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
