package main
import(
	//"fmt"
	//"golang.org/x/net/html"
	//"net/http"
	"debugger"
        //"razor"


	//"os"
	"razor"
	//"scheduler"

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


	razor.GetLiveStockData_GoogleAPI()
//----------------------------------------------------------------------
	//scheduler.NewSchedule("GetLiveStockFromGoogle")
	//scheduler.NewJob("get stock every 1min","0 0/1 * ? 1-12 1-5",razor.GetLiveStockData_GoogleAPI)

	//test---------------------------------------------------------------------
	//scheduler.NewJob("1s","@every 1s",func() { fmt.Println("Every 1s") })
        //scheduler.NewJob("5s","@every 5s",func() { fmt.Println("Every 5s") })
	// test------------------------------------------------------------------

	//scheduler.ScheduleStart()




}



