package helper

import (
	"app/model"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

var instance *gorm.DB

func GetHistories() {

	// Fetch data from API
	data, err := fetchData()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate change percentage
	calculateChange(&data)

	// Save data to database
	err = saveData(&data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data saved successfully")
}

func fetchData() ([]model.OHLCData, error) {
	url := "https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=1&precision=1"
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

	var ohlcData []model.OHLCData
	for _, d := range data {
		ohlc := model.OHLCData{
			ID:        "bitcoin",
			Timestamp: int64(d[0].(float64) / 1000), // Convert milliseconds to seconds
			Open:      d[1].(float64),
			High:      d[2].(float64),
			Low:       d[3].(float64),
			Close:     d[4].(float64),
		}
		ohlcData = append(ohlcData, ohlc)
	}
	return ohlcData, nil
}

func calculateChange(data *[]model.OHLCData) {
	for i := range *data {
		if i > 0 {
			previousClose := (*data)[i-1].Close
			(*data)[i].Change = (((*data)[i].Close - previousClose) / previousClose) * 100
		} else {
			(*data)[i].Change = 0
		}
	}
}

func saveData(data *[]model.OHLCData) error {
	for _, d := range *data {
		// Check if the record already exists
		var count int64
		instance.Model(&model.OHLCData{}).Where("id = ? AND timestamp = ?", d.ID, d.Timestamp).Count(&count)
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
