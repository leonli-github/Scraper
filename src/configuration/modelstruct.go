package configuration


type Modelstruct1 struct {
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64  `json:"low"`
	Close   float64  `json:"close"`
	Volume   float64  `json:"volume"`
	MACD   float64  `json:"macd"`
	MACDSignal   float64  `json:"macdsignal"`
	RSI   float64  `json:"rsi"`
	Label int `json:"label"`
}
