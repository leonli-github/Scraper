package ta
import(
	c "configuration"
        talib "github.com/d4l3k/talib"
	"io/ioutil"
	"encoding/json"
	//"runtime"
	"debugger"
	"os"
	"fmt"
	"strconv"

	"bufio"
)

func GenerateModel(){
	f, _ := os.Create("test.log")
	defer f.Close()
	w := bufio.NewWriter(f)

	for _,stock := range c.Config.StockCode{
          file:= "data/" + stock + "_HK.json"
		_, err := os.Stat(file)
		if err != nil {
			panic("file "+ file + "not exist")
		}
		PrepareModel(w,file)

	}
	fmt.Println("Done")
}


func PrepareModel(w *bufio.Writer, file string){

	historical :=  ImportHistorialdata(file)
	dataset :=  LoadDataSet("close", historical)


	rsi:=CalcuRSI(dataset)
	rsi = shift(rsi,14)
	macd,macdsig := CalcuMACD(dataset)
	macd = shift(macd,33)
	macdsig = shift(macdsig,33)

	var models = make([]c.Modelstruct1,len(*historical))




	for index,element := range *historical{

		models[index].High = element.High
		models[index].Low = element.Low
		models[index].Open = element.Open
		models[index].Close = element.Close
		models[index].Volume = element.Volume
		models[index].RSI = rsi[index]
		models[index].MACD = macd[index]
		models[index].MACDSignal = macdsig[index]

		//fmt.Println(models[index].Open,models[index].Close,models[index].RSI,models[index].MACD,models[index].MACDSignal)

	}



	for i:= 33; i<len(models)-7;i++ {
		trend1 := models[i+5].Close-models[i+5].Open
		trend2 := models[i+6].Close-models[i+6].Open

		label:=0

		if trend1>0 && trend2>0 {
			label = 1
		}
		if trend1<0 && trend2<0 {
			label = -1
		}

		var input string
		for j:=0; j<=4 ; j++{
			input +=  FloatToString(models[i+j].High) +" "
			input +=  FloatToString(models[i+j].Low) +" "
			input +=  FloatToString(models[i+j].Open) +" "
			input +=  FloatToString(models[i+j].Close) +" "
			input +=  FloatToString(models[i+j].High) +" "
			input +=  FloatToString(models[i+j].Volume) +" "
			input +=  FloatToString(models[i+j].RSI) +" "
			input +=  FloatToString(models[i+j].MACD) +" "
			input +=  FloatToString(models[i+j].MACDSignal) +" "
		}
		input += strconv.Itoa(label) + "\n"

		fmt.Println(input)



		_, err := w.WriteString(input)
		if err !=nil {
			fmt.Println(err)
		}
		w.Flush()

	}

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 3, 64)
}

func ImportHistorialdata(file string) *[]c.Diurnal{
	historical := &[]c.Diurnal{}
	bytes, err := ioutil.ReadFile(file)
	if err == nil {

		json.Unmarshal(bytes,historical)
		//fmt.Println(historical)
		// CalcuRSI(data)

	}else{
		debugger.Log(err)
	}
	return  historical

}

func LoadDataSet(field string, historical *[]c.Diurnal) []float64{

	var data []float64 = make([]float64,len(*historical))

	if field == "close"{
		for index,element := range *historical{
                     data[index] = element.Close
			//fmt.Println(element.Date,data[index])
		}
	}

	//data2:= []float64{44.34,44.09,44.15,43.61,44.33,44.83,45.10,45.42,45.84,46.08,45.89,46.03,45.61,46.28,46.28,46.00,46.03,46.41,46.22}

	//rsi := CalcuRSI(data2)
        //rsi = shift(rsi,14)
	//for index,element := range data2{
	//	fmt.Println(element,rsi[index])
	//}
	return data
}

func CalcuRSI(data []float64) []float64{

	rsi := talib.Rsi(data,14)
	return rsi

}


func CalcuMACD(data []float64) ([]float64,[]float64){


        MACD,MACDSignal,_ := talib.Macd(data,12,26,9)
        return MACD,MACDSignal

}

func reverse(numbers []float64) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func shift(arr []float64, number int) []float64{


	l := len(arr)

	part1 := arr[0:l-number]
	part2 := arr[l-number:]

	part2 = append(part2,part1...)

	return part2

}