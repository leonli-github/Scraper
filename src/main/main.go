package main
import(
	//"fmt"
	//"golang.org/x/net/html"
	//"net/http"
	"debugger"
        //"razor"


	//"os"
	"razor"
	"scheduler"

	//"fmt"
)

func main() {



	//get flag
	//flag := os.Args[1]
	//fmt.Println(flag)
        //

	//*initiate debugger*
	file := debugger.Logfile
	defer file.Close()
	//*initiate debugger*


	scheduler.NewSchedule("GetLiveStockFromGoogle")
	//test
	scheduler.NewJob("get stock every 1min","0 0/1 * ? 1-12 1-5",razor.GetLiveStockData_GoogleAPI)

	//scheduler.NewJob("1s","@every 1s",func() { fmt.Println("Every 1s") })
        //scheduler.NewJob("5s","@every 5s",func() { fmt.Println("Every 5s") })
   
	scheduler.ScheduleStart()

	//c.AddFunc("@every 10s", func() { fmt.Println("Every 10s") })


}



