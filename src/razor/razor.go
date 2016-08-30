package razor
import(
	"debugger"
	//"fmt"
	//"golang.org/x/net/html"
	"net/http"
	//"os"
	//"strings"
	//"net/url"
	//"io/ioutil"
	"io/ioutil"
	//"encoding/csv"
)


const  Tempurl  = "http://www.aastocks.com/en/LTP/RTQuote.aspx?symbol=00911"

func Scraper(){
	debugger.Log("razor test3")
}

func Httprequest(url string){
	resp,_ :=http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	debugger.Log(string(bytes))
	resp.Body.Close()
	
}

func parseCSV(){

}