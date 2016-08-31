package scheduler

import(
	"github.com/robfig/cron"
	//"razor"
	"fmt"
)
var ScheduleMap map[string]*cron.Cron

func ini(){

	ScheduleMap = make(map[string]*cron.Cron)
}

func NewJob(){

	c := cron.New()
	fmt.Println("start to add func")
	c.AddFunc("@every 10s", func() { fmt.Println("Every 10s") })
	c.AddFunc("@every 1s", func() { fmt.Println("Every 1s") })
	fmt.Println("c start")
	c.Start()
	defer c.Stop()
	select {

	}

}
