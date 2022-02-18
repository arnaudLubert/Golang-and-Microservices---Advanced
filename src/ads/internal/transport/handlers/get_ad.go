package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAdHandler(cmd utils.GetAdCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ad, err := cmd(req.Context(), mux.Vars(req)["ad_id"])

		if err != nil {
			switch err {
			case domain.ErrAdNotFound:
				http.Error(rw, err.Error(), http.StatusNotFound)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		} else {
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(ad)
		}
	}
}
