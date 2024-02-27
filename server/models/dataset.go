package models

import (
	"encoding/json"
)

type Weather struct {
	Date          string  `json:"date"`
	Temperature   float64 `json:"temperature"`
	Humidity      float64 `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	Precipitation float64 `json:"precipitation"`
}
type BirthRate struct {
	Year      int     `json:"year"`
	BirthRate float64 `json:"birth_rate"`
}

type Internet struct {
	Location      string  `json:"location"`
	InternetSpeed float64 `json:"internet_speed"`
}

func FetchDatasetsFromSource(originalData string, v interface{}) error {
	return json.Unmarshal([]byte(originalData), v)
}
