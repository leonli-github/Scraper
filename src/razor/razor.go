package razor
import(

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

	//"time"
	"debugger"
)


const  Tempurl  = "https://www.etnet.com.hk/www/eng/stocks/realtime/quote.php?code=00911"

func GetLiveStockData_GoogleAPI(){

	url := configuration.GOOGLE_LONG_REAL_TIME_URL + "2388,0911"
	jsonbytes := Httprequest(url)
	//fmt.Println(string(jsonbytes))
	stocks := []configuration.GoogleLongStockLiveStruct{}
	error := json.Unmarshal(jsonbytes[3:],&stocks)
	//fmt.Println(error)
	if error == nil {
		debugger.Log(stocks)
		fmt.Println(stocks[0])
		//fmt.Println(stocks[1])

	}


}
func Httprequest(url string) []byte{
	resp,_ :=http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	//debugger.Log(string(bytes))
	defer resp.Body.Close()
        return bytes
}

func Rz(){

	bytes := Httprequest(Tempurl)
	debugger.Log(string(bytes))

}