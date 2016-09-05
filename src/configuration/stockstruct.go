package configuration

type GoogleShortStockLiveStruct struct {
      Id string `json:"id"`
      Ticker string `json:"t"`
      Exchange string 	`json:"e"`
      Lastprice string `json:"l"`
      Lasttradedatetime string `json:"lt_dts"`
      Change string `json:"c"`
      Changepercentage string `json:"cp"`
}


type GoogleLongStockLiveStruct struct {
      Id string `json:"id"`
      Ticker string `json:"t"`
      Exchange string 	`json:"e"`
      Lastprice string `json:"l"`
      Lasttradedatetime string `json:"lt_dts"`
      Change string `json:"c"`
      Changepercentage string `json:"cp"`
      High string `json:"hi"`
      Low string `json:"lo"`
      Volumn string `json:"vo"`
      Averagevolumn string `json:"avvo"`
      Marketcapital string `json:"mc"`
      PE string `json:"pe"`
      EPS string `json:"eps"`
      Shares string `json:"shares"`
}

type StockTickerTapeStruct struct {
      Id string `json:"id"`
      Timestamp string `json:"timestamp"`
      Ticker string `json:"ticker"`
      Values []GoogleLongStockLiveStruct `json:"values"`
}