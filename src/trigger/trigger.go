package trigger

import(

	ta "github.com/d4l3k/talib"
	c "configuration"
)



func FeedLiveDataToEngine(tm map[string]c.StockTickerTapeStruct){
	tickermap := tm
	Processing()
}

func Processing(){



}

func GenerateStrategy(tickermap map[string]c.StockTickerTapeStruct) c.Strategy{

	strategy := &c.Strategy{}
	
	return strategy
}



