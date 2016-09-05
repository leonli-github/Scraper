package scheduler

import(
	"github.com/robfig/cron"
	//"razor"
	//"fmt"
	"debugger"
)
var cr *cron.Cron
var ScheduleName string

func NewSchedule(scheName string){
    ScheduleName = scheName
    cr = cron.New()
    debugger.Log("Schedule:-" +ScheduleName +"-initiated" )

}

func NewJob(jobName string, timeExp string, f func()){

	//c := cron.New()
	//fmt.Println("start to add func")
	//c.AddFunc("@every 10s"git cannot push github/robfig.cron, func() { fmt.Println("Every 10s") })
	//c.AddFunc("@every 1s", func() { fmt.Println("Every 1s") })
	cr.AddFunc(timeExp,f)
	debugger.Log("Add job: " + jobName)
	//fmt.Println("c start")
	//c.Start()
	//defer c.Stop()
	//select {}

}
func ScheduleStart(){
	cr.Start()
	defer cr.Stop()
	select{}
}
