package middlewares

import (
	"src/ads/domain"
	"src/ads/internal/conf"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func AuthorizeSession(authService conf.Service) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			authRequest, err := http.NewRequest("GET", authService.Url+"/session", nil)

			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			authRequest.Header.Set("x-api-key", authService.ApiKey)
			authRequest.Header.Set("Authorization", req.Header.Get("Authorization"))
			authRequest.Header.Add("Content-Type", "application/json")
			authRequest.Header.Add("Host", authService.Url)

			client := &http.Client{Timeout: time.Second * 20}
			sessionResponse, err := client.Do(authRequest)

			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

			if sessionResponse.StatusCode >= 500 {
				rw.WriteHeader(http.StatusInternalServerError)
				sessionResponse.Body.Close()
				return
			} else if sessionResponse.StatusCode < 200 || sessionResponse.StatusCode > 299 {
				rw.WriteHeader(http.StatusForbidden)
				sessionResponse.Body.Close()
				return
			}
			var sessionInfo domain.SessionInfo

			decoder := json.NewDecoder(sessionResponse.Body)

			if err = decoder.Decode(&sessionInfo); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				sessionResponse.Body.Close()
				return
			}
			sessionResponse.Body.Close()
			ctx := context.WithValue(req.Context(), "session", sessionInfo)
			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}
}
