package domain

import "errors"

var (
	ErrAlreadyExists        = errors.New("Ad already exists")
	ErrAdTitleAlreadyExists = errors.New("Ad title already exists")
	ErrAdNoTitle            = errors.New("you must add a title")
	ErrAdNoDescription      = errors.New("you must add a description")
	ErrAdNoPrice            = errors.New("you must add a price")
	ErrAdNoPicture          = errors.New("you must add a picture")
	ErrInvalidAdValues      = errors.New("Ad values are not valid")
	ErrAdNotFound           = errors.New("Ad not found")
)

type Ad struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
	SellerID    string  `json:"sellerId"`
}
