package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"encoding/json"
	"net/http"
)

func SearchAdsHandler(cmd utils.SearchAdsCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		urlParameters := req.URL.Query()
		ads, err := cmd(req.Context(), urlParameters.Get("sellerId"), urlParameters["keywords"])
		if err != nil {
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
