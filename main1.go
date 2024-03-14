package main

import (
	"app/model"
	"app/payload"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type OHLCData struct {
	ID        string  `json:"-"`
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Change    float64 `json:"change"`
}

var instance *gorm.DB

func main() {

	// Fetch data from API
	data, err := fetchData()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate change percentage
	calculateChange(&data)

	//// Save data to database
	//err = saveData(&data)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Data saved successfully")
}

func fetchData() ([]model.OHLCDataSaveData, error) {

	request := payload.RequestGetHistories{}

	period := request.Period
	period = "MAX"
	var num int
	var unit string
	if period == "MAX" {
		unit = "MAX"
	}
	fmt.Sscanf(period, "%d%s", &num, &unit)

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
	dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	instance = db
	log.Println("Connected to the database")

	id := request.Symbol
	result := db.Raw("SELECT * FROM coins where symbol = ?", id).First()

	if num != 0 && num <= 7 && period == "30M" && unit == "H" {
		days = 1
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
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data [][]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	var ohlcData []model.OHLCDataSaveData
	for _, d := range data {
		// Ensure each element is an array of length 5
		if len(d) != 5 {
			return nil, fmt.Errorf("unexpected data format")
		}

		ohlc := model.OHLCDataSaveData{
			ID:        string(result),
			Timestamp: int64(d[0].(float64) / 1000), // Convert milliseconds to seconds
			Open:      d[1].(float64),
			High:      d[2].(float64),
			Low:       d[3].(float64),
			Close:     d[4].(float64),
		}
		ohlcData = append(ohlcData, ohlc)
	}

	fmt.Println(ohlcData)
	return ohlcData, nil

}

func calculateChange(data *[]model.OHLCDataSaveData) {
	for i := range *data {
		if i > 0 {
			previousClose := (*data)[i-1].Close
			(*data)[i].Change = (((*data)[i].Close - previousClose) / previousClose) * 100
		} else {
			(*data)[i].Change = 0
		}
	}
}

func saveData(data *[]model.OHLCDataSaveData) error {
	for _, d := range *data {
		// Check if the record already exists
		var count int64
		instance.Model(&model.OHLCDataSaveData{}).Where("id = ? AND timestamp = ?", d.ID, d.Timestamp).Count(&count)
		if count == 0 {
			// If the record does not exist, create a new one
			result := instance.Create(&d)
			if result.Error != nil {
				return result.Error
			}
		} else {
			// If the record exists, skip
			log.Printf("Record with id '%s' and timestamp '%d' already exists. Skipping...", d.ID, d.Timestamp)
		}
	}
	return nil
}
