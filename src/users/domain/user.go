package domain

import "errors"

var (
    ErrUserNotFound = errors.New("user not found")
    ErrUserAlreadyExists = errors.New("user already exists")
    ErrUserLoginAlreadyExists = errors.New("another account is using this login")
    ErrUserEmailAlreadyExists = errors.New("another account is using this email address")
)

type User struct {
    ID          string  `json:"id"`
    Email       string  `json:"email"`
    Login       string  `json:"login"`
    Firstname   string  `json:"firstname"`
    Lastname    string  `json:"lastname"`
    Phone       string  `json:"phone"`
    Password    string  `json:"-"`
    Access      int8    `json:"access"`
    Address     Address `json:"address"`
}