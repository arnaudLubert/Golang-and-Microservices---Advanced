package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateAdRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func UpdateAdHandler(cmd utils.UpdateAdCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var updateAdReq UpdateAdRequest

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&updateAdReq)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ad := domain.Ad{
			Title:       updateAdReq.Title,
			Description: updateAdReq.Description,
			Price:       updateAdReq.Price,
		}

		if err = cmd(req.Context(), mux.Vars(req)["ad_id"], ad); err != nil {
			switch err {
			case domain.ErrAdNotFound:
				http.Error(rw, err.Error(), http.StatusNotFound)
			case domain.ErrAlreadyExists:
				http.Error(rw, err.Error(), http.StatusConflict)
			case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
				http.Error(rw, err.Error(), http.StatusForbidden)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		} else {
			rw.WriteHeader(http.StatusNoContent)
		}
	}
}
