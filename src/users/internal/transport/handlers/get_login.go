package handlers

import (
	"encoding/json"
	"net/http"
	"src/users/domain"
	"src/users/internal/security"
	"src/users/internal/utils"
)

type GetLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetLoginResponse struct {
	UserID string `json:"user_id"`
}

func GetLoginHandler(cmd utils.GetLoginCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var getLoginRequest GetLoginRequest

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&getLoginRequest)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		userId, err := cmd(req.Context(), getLoginRequest.Login, security.MD5(getLoginRequest.Password))

		if err != nil {
			switch err {
			case domain.ErrUserNotFound:
				http.Error(rw, err.Error(), http.StatusNotFound)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		} else {
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(GetLoginResponse{UserID: userId})
		}
	}
}
