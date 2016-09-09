package scheduler

import (
	"github.com/robfig/cron"
	//"razor"
	//"fmt"
	"debugger"
)

var cr *cron.Cron
var ScheduleName string

func NewSchedule(scheName string) {
	ScheduleName = scheName
	cr = cron.New()
	debugger.Log("Schedule:-" + ScheduleName + "-initiated")

}

func NewJob(jobName string, timeExp string, f func()) {

	cr.AddFunc(timeExp, f)
	debugger.Log("Add job: " + jobName)

}
func ScheduleStart() {
	cr.Start()
	defer cr.Stop()
	select {}
}
