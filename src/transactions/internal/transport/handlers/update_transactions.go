package handlers

import (
	"encoding/json"
	"net/http"
	"src/transactions/domain"
	"src/transactions/internal/utils"

	"github.com/gorilla/mux"
)

func UpdateHandler(cmd utils.UpdateCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var updateReq domain.Invoice

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&updateReq)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		partialTsn := domain.Invoice{
			Message:  updateReq.Message,
			BidPrice: updateReq.BidPrice,
			AdID:     updateReq.AdID,
		}

		if err = cmd(req.Context(), mux.Vars(req)["id"], partialTsn); err != nil {
			handleErrors(rw, err)
		} else {
			rw.WriteHeader(http.StatusNoContent)
		}
	}
}
