package domain

import  (
    "errors"
    //"time"
)

var (
    ErrReservationAlreadyExists = errors.New("Reservation already exists")
    ErrDateStringInvalid = errors.New("Reservztion date is invalid")
    // ErrUserNotFound = errors.New("user not found")
    // ErrUserAlreadyExists = errors.New("user already exists")
    // ErrUserPseudoAlreadyExists = errors.New("another account is using this pseudo")
    // ErrUserEmailAlreadyExists = errors.New("another account is using this email address")
    // ErrUnsecuredPassword = errors.New("this password is not secured, it must contain 8 characters and at least 1 uppercase, 1 lowercase and 1 digit")
)

// type User struct {
//     ID          string  `json:"id"`
//     Email       string  `json:"email"`
//     Pseudo       string  `json:"pseudo"`
//     Firstname   string  `json:"firstname"`
//     Lastname    string  `json:"lastname"`
//     Phone       string  `json:"phone"`
//     Password    string  `json:"-"`
//     Access      int8    `json:"access"`
//     Address     Address `json:"address"`
//     IBAN        string  `json:"iban"`
// }

type Reservation struct {
    ID                  string      `json:"id"`
    Date_of_arrival     string      `json:"date_begin"` //time.Time
    Number_night        int64       `json:"nb_night"`
    Number_people       int64       `json:"nb_people"`
    Status              Resa_Status `json:"status"`
    Ad_id               string      `json:"ad_id"`
    Owner_id            string      `json:"owner_id"`
    Customer_id         string      `json:"customer_id"`
}

type Resa_Status int64

const (
    pending Resa_Status = 0
    accepted            = 1
    refused             = 2
)