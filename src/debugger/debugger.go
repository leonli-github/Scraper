package debugger

import(
	"log"
	"os"
	//"fmt"
)

var logwriter debugger_s

var Logfile *os.File

type debugger_s struct{
	logger *log.Logger
}


func init() {
	logwriter = debugger_s{}
	Logfile = logwriter.initdebugger_s()
}

func Log(s string) {

	logwriter.logger.Println(s)
}



func (l *debugger_s) initdebugger_s() (logfile *os.File) {
	fileName := "xxx_debug.html"
	logfile,err  := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error !")
	}
       l.logger = log.New(logfile,"[Debug]",log.Llongfile)
         return
}


