package domain

import "errors"

var (
	ErrAlreadyExists        = errors.New("Ad already exists")
	ErrAdTitleAlreadyExists = errors.New("Ad title already exists")
	ErrAdNoTitle            = errors.New("you must add a title")
	ErrAdNoDescription      = errors.New("you must add a description")
	ErrAdNoPrice            = errors.New("you must add a price")
	ErrAdNoCapacity         = errors.New("you must add a capacity")
	ErrAdNoPicture          = errors.New("you must add a picture")
	ErrAdNoLocation         = errors.New("you must add a location")
	ErrInvalidAdValues      = errors.New("Ad values are not valid")
	ErrAdNotFound           = errors.New("Ad not found")
)

type Ad struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Capacity    int       `json:"capacity"`
	Price       float64   `json:"price"`
	Pictures    []string  `json:"pictures"`
	SellerID    string    `json:"sellerId"`
	Location    Location  `json:"location"`
	//Availability    string    `json:"Availability"`
}
