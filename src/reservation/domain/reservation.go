package domain

import "errors"

var (
    ErrUserNotFound = errors.New("user not found")
    ErrUserAlreadyExists = errors.New("user already exists")
    ErrUserPseudoAlreadyExists = errors.New("another account is using this pseudo")
    ErrUserEmailAlreadyExists = errors.New("another account is using this email address")
    ErrUnsecuredPassword = errors.New("this password is not secured, it must contain 8 characters and at least 1 uppercase, 1 lowercase and 1 digit")
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