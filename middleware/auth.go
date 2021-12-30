package middleware

import (
	"app"
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO:
		uid := int64(123456)
		ctx := context.WithValue(r.Context(), app.UserIdKey, uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
