package domain

type Address struct {
    City    string  `json:"city"`
    ZipCode string  `json:"zip_code"`
    Street  string  `json:"street"`
}