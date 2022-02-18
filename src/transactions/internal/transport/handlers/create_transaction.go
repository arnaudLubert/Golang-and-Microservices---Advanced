package handlers

import (
	"encoding/json"
	"net/http"
	"src/transactions/domain"
	"src/transactions/internal/utils"
)

type CreateTsnResponse struct {
	TsnID string `json:"transaction_id"`
}

func CreateHandler(cmd utils.CreateCmd, getAd GetAdCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var createTsnReq domain.Invoice

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&createTsnReq)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ad, err := getAd(createTsnReq.AdID)
		if err != nil {
			handleErrors(rw, err)
			return
		}

		if createTsnReq.BidPrice == 0 {
			handleErrors(rw, domain.ErrRequiredField)
			return
		}
		tsn := domain.Transaction{
			BidPrices:  []float64{createTsnReq.BidPrice},
			AdID:       createTsnReq.AdID,
			AdSellerID: ad.SellerID,
		}
		if createTsnReq.Message != "" {
			tsn.Messages = []string{createTsnReq.Message}
		}
		tsnID, err := cmd(req.Context(), tsn)

		if err != nil {
			handleErrors(rw, err)
		} else {
			rw.WriteHeader(http.StatusCreated)
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(CreateTsnResponse{tsnID})
		}
	}
}
