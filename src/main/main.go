package main
import(
	"fmt"
	//"golang.org/x/net/html"
	//"net/http"
	"debugger"
        //"razor"
	"scheduler"

)

func main() {

    //*initiate debugger*
	file := debugger.Logfile
	defer file.Close()
	//*initiate debugger*

	//razor.Httprequest(razor.Tempurl)
	scheduler.NewSchedule("Test Cron")


	scheduler.NewJob("2s","@every 2s",func() { fmt.Println("Every 2s") })
    scheduler.NewJob("5s","@every 5s",func() { fmt.Println("Every 5s") })
   
     scheduler.ScheduleStart()

	//c.AddFunc("@every 10s", func() { fmt.Println("Every 10s") })


}



