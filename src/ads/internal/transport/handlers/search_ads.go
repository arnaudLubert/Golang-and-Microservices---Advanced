package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func SearchAdsHandler(cmd utils.SearchAdsCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var distance float64
		var err error

		query := req.URL.Query()
		coordinates := make([]float64, 2)

		if coordinates[0], err = strconv.ParseFloat(query.Get("coordinate_latitude"), 64); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		} else if coordinates[1], err = strconv.ParseFloat(query.Get("coordinate_longitude"), 64); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		} else if distance, err = strconv.ParseFloat(query.Get("distance"), 64); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if ads, err := cmd(req.Context(), coordinates, distance); err != nil {
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
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(ads)
		}
	}
}
