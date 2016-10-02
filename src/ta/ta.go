package ta
import(
	c "configuration"
        talib "github.com/d4l3k/talib"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func ImportHistorialdata(){
	historical := &[]c.Diurnal{}
	bytes, err := ioutil.ReadFile("../ta/0001_HK.json")
	if err == nil {

		json.Unmarshal(bytes,historical)

		fmt.Println(historical)

	}

}
