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
	"ta"
	"kafka"
)

const nottest = 2

func main() {

	//*get flag*
	//flag := os.Args[1]
	//fmt.Println(flag)
	//#get flag*
	if nottest == 0 {
		//*initiate debugger*
		file := debugger.Logfile
		defer file.Close()
		//*initiate debugger*

		//razor.GetLiveStockData_GoogleAPI()
		razor.InitiateTickermap()
		//----------------------------------------------------------------------
		scheduler.NewSchedule("GetLiveStockFromGoogle")
		//scheduler.NewJob("get stock every 1min", "0 0/1 * ? 1-12 1-5", razor.GetLiveStockData_GoogleAPI)

		//test---------------------------------------------------------------------
		//scheduler.NewJob("1s","@every 1s",func() { fmt.Println("Every 1s") })
		scheduler.NewJob("5s", "@every 5s", razor.GetLiveStockData_GoogleAPI)
		// test------------------------------------------------------------------

		scheduler.ScheduleStart()
	} else if nottest == 1 {
		ta.GenerateModel()
	} else if nottest == 2 {
		kafka.ProducerMain()
	}

}



