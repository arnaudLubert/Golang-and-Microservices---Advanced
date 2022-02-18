package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"net/http"
)

func DeleteAdHandler(cmd utils.DeleteAdCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		err := cmd(req.Context(), req.URL.Query().Get("ad_id"))

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
			rw.WriteHeader(http.StatusNoContent)
		}
	}
}
