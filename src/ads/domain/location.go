package domain

type Location struct {
    City       string  `json:"city"`
    ZipCode    string  `json:"zip_code"`
    Street     string  `json:"street"`
    Longitude  string  `json:"longitude"`
    Latitude   string  `json:"latitude"`
}