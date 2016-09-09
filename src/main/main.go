package main

import (
	//"fmt"
	//"golang.org/x/net/html"
	//"net/http"
	"debugger"
	//"os"
	//"razor"
	//"scheduler"
	//"math"
	//"github.com/d4l3k/talib"
	"scheduler"
	"razor"
	//"math"
	//"fmt"
)

func main() {

	//*get flag*
	//flag := os.Args[1]
	//fmt.Println(flag)
	//#get flag*

	//*initiate debugger*
	file := debugger.Logfile
	defer file.Close()
	//*initiate debugger*

	//fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))

	//razor.GetLiveStockData_GoogleAPI()
	razor.InitiateTickermap()
	//razor.Rz()
	//----------------------------------------------------------------------
	scheduler.NewSchedule("GetLiveStockFromGoogle")
	scheduler.NewJob("get stock every 1min", "0 0/1 * ? 1-12 1-5", razor.GetLiveStockData_GoogleAPI)

	//test---------------------------------------------------------------------
	//scheduler.NewJob("1s","@every 1s",func() { fmt.Println("Every 1s") })
	//scheduler.NewJob("1m","@every 1min",razor.GetLiveStockData_GoogleAPI)
	// test------------------------------------------------------------------

	scheduler.ScheduleStart()

}



