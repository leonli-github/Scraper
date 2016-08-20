package debugger

import(
	"log"
	"os"
)

var logwriter debugger_s

type debugger_s struct{
	logger *log.Logger
}


func Init() (logfile *os.File){
	logwriter = debugger_s{}
	logfile = logwriter.init()
	return
}
func Log(s string) {

	logwriter.logger.Println(s)
}


func (l *debugger_s) init() (logfile *os.File) {
	fileName := "xxx_debug.log"
	logfile,err  := os.Create(fileName)
	if err != nil {
		log.Fatalln("open file error !")
	}
       l.logger = log.New(logfile,"[Debug]",log.Llongfile)
         return
}


