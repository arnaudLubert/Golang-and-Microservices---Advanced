package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"src/transactions/internal/utils"
)

func ActionHandler(cmd utils.ActionCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		err := cmd(req.Context(), mux.Vars(req)["id"])

		if err != nil {
			handleErrors(rw, err)
		} else {
			rw.WriteHeader(http.StatusNoContent)
		}
	}
}
