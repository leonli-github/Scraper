package configuration

import "time"

type Modelstruct1 struct {
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64  `json:"low"`
	Volume   float64  `json:"volume"`
	MACD   float64  `json:"macd"`
	RSI   float64  `json:"rsi"`
	Date  time.Time `json:"date"`
}
