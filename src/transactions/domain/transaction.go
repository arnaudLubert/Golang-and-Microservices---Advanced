package domain

import "errors"

var (
	ErrNotFound             = errors.New("transaction not found")
	ErrAlreadyExists        = errors.New("transaction already exists")
	ErrRequiredField        = errors.New("one or multiple field are missing")
	ErrCannotRetrieveSeller = errors.New("cannot retrieve ad owner")
)

type Invoice struct {
	Message  string  `json:"message"`
	BidPrice float64 `json:"bid_price"`
	AdID     string  `json:"ad_id"`
}

type Transaction struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	BidPrices  []float64 `json:"bid_prices"`
	Messages   []string  `json:"messages"`
	SenderID   string    `json:"sender_id"`
	AdID       string    `json:"ad_id"`
	AdSellerID string    `json:"seller_id"`
}
