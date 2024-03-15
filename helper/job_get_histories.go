package helper

import (
	"app/dbutil"
	"app/model"
	"app/payload"
	"app/usecases"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func FetchData(period string, symbol string) {

	var num int
	var unit string
	if period == "MAX" {
		unit = "MAX"
	}
	fmt.Sscanf(period, "%d%s", &num, &unit)
	fmt.Println("period: ", period)

	var days int
	var url string

	//Trường hợp nhỏ hơn 14 day
	b := 7 < num && num <= 14 && unit == "D"

	//Trường hợp nhỏ hơn 30 day
	c := 14 < num && num <= 30 && unit == "D"

	//Trường hợp nhỏ hơn 90 day
	d := 30 < num && num <= 90 && unit == "D"

	//Trường hợp nhỏ hơn 180 day
	e := 90 < num && num <= 180 && unit == "D"

	//Trường hợp nhỏ hơn 365 day
	f := 180 < num && num <= 365 && unit == "D"
	db := dbutil.ConnectDB()

	id := strings.ToLower(symbol)

	request := payload.RequestGetHistories{}

	period = request.Period
	fmt.Println("period: ", period)

	var result model.OHLCDataSaveData
	err := db.Raw("SELECT * FROM coins WHERE symbol = ? LIMIT 1", id).Scan(&result).Error
	if err != nil {
		log.Fatal(err)
	}
	// In giá trị của dòng đầu tiên ra
	//fmt.Printf("ID: %s\n", result.ID)

	if num == 1 && num <= 6 || unit == "H" || period == "30M" {
		days = 1
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if num <= 7 && unit == "D" {
		days = 7
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if b {
		days = 14
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if c {
		days = 30
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if d {
		days = 90
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if e {
		days = 180
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if f {
		days = 365
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%d", days)
	} else if num == 0 && period == "MAX" {
		unit = strings.ToLower(unit)
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=%s", unit)
	}

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error get url:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error get body:", err)
	}

	var data [][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error :", err)
	}

	var ohlcData []model.OHLCDataSaveData
	for _, d := range data {
		// Ensure each element is an array of length 5
		if len(d) != 5 {
			log.Println("Error data:", err)
		}

		ohlc := model.OHLCDataSaveData{
			ID:        result.ID,
			Timestamp: int64(d[0].(float64) / 1000), // Convert milliseconds to seconds
			Open:      d[1].(float64),
			High:      d[2].(float64),
			Low:       d[3].(float64),
			Close:     d[4].(float64),
		}
		ohlcData = append(ohlcData, ohlc)
	}
	// Declare a slice instead of a pointer to slice
	var dataa []model.OHLCDataSaveData

	// Populate the slice with the fetched data
	for _, ohlc := range ohlcData {
		dataa = append(dataa, ohlc)
	}

	// Calculate change for each data point
	for i := range dataa {
		if i > 0 {
			previousClose := dataa[i-1].Close
			dataa[i].Change = (((dataa[i].Close - previousClose) / previousClose) * 100)
		} else {
			dataa[i].Change = 0
		}
	}

	// Print the modified dataa slice
	//fmt.Println(dataa)

	uc := usecases.SaveOhlcDataUseCase()
	if err := uc.SaveOhlcData(dataa); err != nil {
		log.Println("Error inserting coins:", err)
	}

}
