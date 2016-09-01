package scheduler

import(
	"github.com/robfig/cron"
	//"razor"
	"fmt"
	"debugger"
)
var c *cron.Cron
var ScheduleName string

func NewSchedule(scheName string){
    ScheduleName = scheName
    c = cron.New()
    debugger.Log("Schedule:-" +ScheduleName +"-initiated" )

}

func NewJob(jobName string, timeExp string, f func()){

	//c := cron.New()
	//fmt.Println("start to add func")
	//c.AddFunc("@every 10s"git cannot push github/robfig.cron, func() { fmt.Println("Every 10s") })
	//c.AddFunc("@every 1s", func() { fmt.Println("Every 1s") })
	c.AddFunc(timeExp,f)
        fmt.Println("Add job: " + jobName)
	//fmt.Println("c start")
	//c.Start()
	//defer c.Stop()
	//select {}

}
func ScheduleStart(){
	c.Start()
	defer c.Stop()
	select{}
}
