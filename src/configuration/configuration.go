package configuration

import (
	"io/ioutil"
	"encoding/json"
	"debugger"
)

//const GOOGLE_SHORT_REAL_TIME_URL = "http://www.google.com/finance/info?client=ig&q="
const GOOGLE_LONG_REAL_TIME_URL = "http://www.google.com/finance/info?infotype=infoquoteall&q="

var Config *Con

type Con struct {
	StockCode []string `json:"stockcode"`
}

func init() {
	Config = &Con{}
	bytes, err := ioutil.ReadFile("../../config/config.json")
	if err == nil {

		json.Unmarshal(bytes, Config)
		debugger.Log("Configuraion file loaded")

	}

}
