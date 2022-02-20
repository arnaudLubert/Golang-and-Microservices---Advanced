package handlers

import (
	"net/http"
	"src/users/domain"
	"src/users/internal/utils"

	"github.com/gorilla/mux"
)

func GetUserIbanHandler(cmd utils.GetUserCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		user, err := cmd(req.Context(), mux.Vars(req)["user_id"])

		if err != nil {
			switch err {
			case domain.ErrUserNotFound:
				http.Error(rw, err.Error(), http.StatusNotFound)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		} else {
			rw.Header().Set("Content-type", "text/plain")
			rw.Write([]byte(user.IBAN))
		}
	}
}
