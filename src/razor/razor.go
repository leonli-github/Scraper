package razor
import(

	//"fmt"
	//"golang.org/x/net/html"
	"net/http"
	//"os"
	//"strings"
	//"net/url"
	//"io/ioutil"
	"io/ioutil"
	"encoding/json"
	c "configuration"
	"debugger"
	"time"
	"strings"
)


//const  Tempurl  = "http://www.aastocks.com/en/ltp/RTQuote.aspx?S=Y&Symbol=00911"

var Tickermap map[string]c.StockTickerTapeStruct
var stockcodestring string


func GetLiveStockData_GoogleAPI(){

	url := c.GOOGLE_LONG_REAL_TIME_URL + stockcodestring
	jsonbytes := Httprequest(url)
	//fmt.Println(string(jsonbytes))
	stocks := []c.GoogleLongStockLiveStruct{}
	error := json.Unmarshal(jsonbytes[3:],&stocks)
	//fmtgolang date.Println(error)
	if error == nil {
		//debugger.Log(stocks)

		for _, value := range stocks{
			Tickermap[value.Ticker] = UpdateTicker(value,Tickermap[value.Ticker])
			debugger.Log(Tickermap[value.Ticker])
		}


	}

}
func Httprequest(url string) []byte{
	resp,_ :=http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	//debugger.Log(string(bytes))
	defer resp.Body.Close()
        return bytes
}

func InitiateTickermap(){
	if Tickermap==nil {
		Tickermap = make(map[string]c.StockTickerTapeStruct)
		debugger.Log("initiated Tickermap")
	}

	for _, value := range c.Config.StockCode {
		tickertape :=new(c.StockTickerTapeStruct)
		tickertape.Ticker = value
		tickertape.Id = "Test"
		tickertape.Timestamp = time.Now().String()
		tickertape.Values = make([]c.GoogleLongStockLiveStruct,0)
		Tickermap[value] = *tickertape
		debugger.Log("Added:"+value+" to Tickermap")

	}

	stockcodestring = strings.Join(c.Config.StockCode,",")
	debugger.Log("Request stock code are: " + stockcodestring)

}

func UpdateTicker(value c.GoogleLongStockLiveStruct,origin c.StockTickerTapeStruct ) c.StockTickerTapeStruct{
	//Tickermap["2388"] .Values = append(tickertape.Values,value)
	origin.Values = append(origin.Values,value)
	return  origin
}

func Rz(){
        //println(c.Config.StockCode)
	//bytes := Httprequest(Tempurl)
	debugger.Log(c.Config.StockCode[0])

}
