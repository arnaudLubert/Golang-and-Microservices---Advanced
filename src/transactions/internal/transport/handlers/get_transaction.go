package handlers

import (
	"encoding/json"
	"net/http"
	"src/transactions/internal/utils"

	"github.com/gorilla/mux"
)

func GetHandler(cmd utils.GetCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tsn, err := cmd(req.Context(), mux.Vars(req)["id"])

		if err != nil {
			handleErrors(rw, err)
		} else {
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(tsn)
		}
	}
}
