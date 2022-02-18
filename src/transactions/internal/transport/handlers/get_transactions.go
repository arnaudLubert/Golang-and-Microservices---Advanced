package handlers

import (
	"encoding/json"
	"net/http"
	"src/transactions/internal/utils"
)

func GetAllHandler(cmd utils.GetAllCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tsnx, err := cmd(req.Context())
		if err != nil {
			handleErrors(rw, err)
		} else {
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(tsnx)
		}
	}
}
