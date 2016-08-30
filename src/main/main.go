package main
import(
	//"fmt"
	//"golang.org/x/net/html"
	//"net/http"
	"debugger"
        "razor"

)

func main() {

        //initiate debugger

	file := debugger.Logfile
	defer file.Close()
	razor.Httprequest(razor.Tempurl)


}



