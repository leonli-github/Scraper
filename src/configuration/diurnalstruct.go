package configuration

import "time"

type Diurnal struct{
	Date  time.Time `json:"date"`
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64  `json:"low"`
	Volume   float64  `json:"volume"`
	Adjclose   float64  `json:"adjClose"`
	Symbol   string  `json:"symbol"`

}
