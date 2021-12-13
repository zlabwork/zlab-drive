package middleware

import (
	"github.com/google/uuid"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.Header.Get("Request-Id")
		if id == "" {
			id = uuid.New().String()
			w.Header().Set("Request-Id", id)
		}

		// Do stuff here
		log.Println(id, r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
