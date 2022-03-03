package handlers

import (
    "src/reservation/internal/utils"
    "src/reservation/domain"
    "encoding/json"
    "net/http"
)


type CreateReservationRequest struct {
    ID                  string      `json:"id"`
    Date_of_arrival     string      `json:"date_begin"` //time.Time
    Number_night        int64       `json:"nb_night"`
    Number_people       int64       `json:"nb_people"`
    Status              Resa_Status `json:"status"`
    Ad_id               string      `json:"ad_id"`
    Owner_id            string      `json:"owner_id"`
    Customer_id         string      `json:"customer_id"`
}

type CreateUserResponse struct {
    ResaID      string          `json:"reservation_id"`
}


func CreateReservationHandler(cmd utils.CreateReservationCmd) http.HandlerFunc {
    return func(rw http.ResponseWriter, req *http.Request) {
        var createReservationReq CreateReservationRequest

        decoder := json.NewDecoder(req.Body)
        err := decoder.Decode(&createReservationReq)

        if err != nil {
            http.Error(rw, err.Error(), http.StatusBadRequest)
            return
        }

        if createReservationReq.Date_of_arrival == "" || createReservationReq.Number_night == "" || createReservationReq.Number_people == "" {
            http.Error(rw, "Reservation missing field", http.StatusBadRequest)
            return
        }

        Reservation := domain.Reservation{
			Date_of_arrival:	createReservationReq.Date_of_arrival,
			Number_night:		createReservationReq.Number_night,
			Number_people:		createReservationReq.Number_people,
			Status: 			0, // pending
			Ad_id: 				createReservationReq.Ad_id,
			Owner_id: 			createReservationReq.Owner_id,
			Customer_id:		createReservationReq.Customer_id,


        }
        reservationID, err := cmd(req.Context(), reservation)

        if err != nil {
			http.Error(rw, err.Error(), http.StatusForbidden)
            // switch err {
            // case domain.ErrUserAlreadyExists:
            //     http.Error(rw, err.Error(), http.StatusConflict)
            // case domain.ErrUnsecuredPassword:
            //     http.Error(rw, err.Error(), http.StatusBadRequest)
            // case domain.ErrAccessForbidden, domain.ErrOperationNotPermitted:
            //     http.Error(rw, err.Error(), http.StatusForbidden)
            // default: http.Error(rw, err.Error(), http.StatusInternalServerError)
            // }
        } else {
            rw.WriteHeader(http.StatusCreated)
            rw.Header().Set("Content-type", "application/json")
            json.NewEncoder(rw).Encode(CreateUserResponse{ResaID: reservationID})
        }
    }
}
