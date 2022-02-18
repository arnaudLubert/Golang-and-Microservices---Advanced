package domain

import "errors"

var (
	ErrAdNotFound = errors.New("ad not found")
)

type Ad struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
	SellerID    string  `json:"sellerId"`
}
