package debugger

import (
	"log"
	"os"
	"fmt"
)

var logwriter debugger_s

var Logfile *os.File

type debugger_s struct {
	logger *log.Logger
}

func init() {
	logwriter = debugger_s{}
	Logfile = logwriter.initdebugger_s()
}

func Log(a interface{}) {

	fmt.Println(a)
	logwriter.logger.Println(a)

}

func (l *debugger_s) initdebugger_s() (logfile *os.File) {

	//production
	//t := time.Now()
	//fileName := t.String() + "_debug.log"
	//-----------------
	fileName := "debug.log"

	logfile, err := os.Create(fileName)

	if err != nil {
		log.Fatalln("open file error !")
	}
	l.logger = log.New(logfile, "[Debug]", log.Ldate | log.Ltime)
	return
}


