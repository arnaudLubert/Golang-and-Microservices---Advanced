package domain

type Location struct {
    City       string  `json:"city"`
    ZipCode    string  `json:"zip_code"`
    Street     string  `json:"street"`
    Latitude   float64 `json:"latitude"`
    Longitude  float64 `json:"longitude"`
}