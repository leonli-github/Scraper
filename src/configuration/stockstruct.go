package configuration

type GoogleStockLiveStruct struct {
      Id string `json:"id"`
      Ticker string `json:"t"`
      Exchange string 	`json:"e"`
      Lastprice string `json:"l"`
      Lasttradedatetime string `json:"lt_dts"`
      Change string `json:"c"`
      Changepercentage string `json:"cp"`
}
