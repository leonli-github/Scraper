package razor
import(
	//"debugger"
	"fmt"
	//"golang.org/x/net/html"
	"net/http"
	//"os"
	//"strings"
	//"net/url"
	//"io/ioutil"
	"io/ioutil"
	"encoding/json"
	"configuration"

	"time"
)


const  Tempurl  = "http://www.aastocks.com/en/LTP/RTQuote.aspx?symbol=00911"

func GetLiveStockData_GoogleAPI(){

	url := configuration.GOOGLE_REAL_TIME_URL + "2388,0911"
	jsonbytes := Httprequest(url)
	//fmt.Println(string(jsonbytes))
	stocks := []configuration.GoogleStockLiveStruct{}
	error := json.Unmarshal(jsonbytes[3:],&stocks)
	//fmt.Println(error)
	if error == nil {
		fmt.Println(stocks[0])
		fmt.Println(stocks[1])
		fmt.Println(time.Now())
	}


}
func Httprequest(url string) []byte{
	resp,_ :=http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	//debugger.Log(string(bytes))
	defer resp.Body.Close()
        return bytes
}

func parseCSV(){

}