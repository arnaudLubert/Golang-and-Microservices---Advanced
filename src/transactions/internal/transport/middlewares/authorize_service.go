package middlewares

import (
	"net/http"
	"src/transactions/internal/conf"

	"github.com/gorilla/mux"
)

func AuthorizeService(credentials conf.Credentials) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			apiKey := req.Header.Get("x-api-key")

			if apiKey != credentials.ApiKey {
				http.Error(rw, "invalid api key", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(rw, req)
		})
	}
}
