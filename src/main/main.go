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

	file := debugger.Init()
	defer file.Close()
	razor.Ri()


}



