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
	"time"
)


const  Tempurl  = "https://www.etnet.com.hk/www/eng/stocks/realtime/quote.php?code=00911"

var Tickermap map[string]configuration.StockTickerTapeStruct

func GetLiveStockData_GoogleAPI(){

	url := configuration.GOOGLE_LONG_REAL_TIME_URL + "2388,0911"
	jsonbytes := Httprequest(url)
	//fmt.Println(string(jsonbytes))
	stocks := []configuration.GoogleLongStockLiveStruct{}
	error := json.Unmarshal(jsonbytes[3:],&stocks)
	//fmt.Println(error)
	if error == nil {
		//debugger.Log(stocks)
		//fmt.Println(stocks[0])
		//fmt.Println(stocks[1])

		Tickermap["2388"] = UpdateTicker(stocks[0],Tickermap["2388"])
		fmt.Println(Tickermap["2388"])
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
		Tickermap = make(map[string]configuration.StockTickerTapeStruct)
		println("initiated Tickermap")
	}

	tickertape :=new(configuration.StockTickerTapeStruct)
	tickertape.Ticker = "2388"
	tickertape.Id = "Test"
	tickertape.Timestamp = time.Now().String()
	tickertape.Values = make([]configuration.GoogleLongStockLiveStruct,0)
	Tickermap["2388"] = *tickertape
	println("stub map")

}

func UpdateTicker(value configuration.GoogleLongStockLiveStruct,origin configuration.StockTickerTapeStruct ) configuration.StockTickerTapeStruct{
	//Tickermap["2388"] .Values = append(tickertape.Values,value)
	origin.Values = append(origin.Values,value)
	return  origin
}

func Rz(){

	bytes := Httprequest(Tempurl)
	debugger.Log(string(bytes))

}
/*
func Append(slice []configuration.GoogleLongStockLiveStruct, data configuration.GoogleLongStockLiveStruct) []configuration.GoogleLongStockLiveStruct {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]configuration.GoogleLongStockLiveStruct, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
*/