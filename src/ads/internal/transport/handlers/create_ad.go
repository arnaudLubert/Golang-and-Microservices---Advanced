package handlers

import (
	"src/ads/domain"
	"src/ads/internal/utils"
	"encoding/json"
	"net/http"
)

type CreateAdRequest struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Capacity    int             `json:"capacity"`
	Price       float64         `json:"price"`
	Pictures    []string        `json:"pictures"`
	Location    domain.Location `json:"location"`
}

type CreateAdResponse struct {
	AdID string `json:"ad_id"`
}

func CreateAdHandler(cmd utils.CreateAdCmd) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		var createAdReq CreateAdRequest

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&createAdReq)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ad := domain.Ad{
			Title:       createAdReq.Title,
			Description: createAdReq.Description,
			Capacity:    createAdReq.Capacity,
			Price:       createAdReq.Price,
			Pictures:    createAdReq.Pictures,
			Location:    createAdReq.Location,
		}
		adID, err := cmd(req.Context(), ad)

		if err != nil {
			switch err {
			case domain.ErrAlreadyExists:
				http.Error(rw, err.Error(), http.StatusConflict)
			case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
				http.Error(rw, err.Error(), http.StatusForbidden)
			case domain.ErrMissingIBAN, domain.ErrAdNoTitle, domain.ErrAdNoDescription, domain.ErrAdNoPrice, domain.ErrAdNoCapacity, domain.ErrAdNoPicture, domain.ErrAdNoLocation:
				http.Error(rw, err.Error(), http.StatusBadRequest)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		} else {
			rw.WriteHeader(http.StatusCreated)
			rw.Header().Set("Content-type", "application/json")
			json.NewEncoder(rw).Encode(CreateAdResponse{AdID: adID})
		}
	}
}
