package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type OHLCData struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"0"`
	Open      float64 `json:"1"`
	High      float64 `json:"2"`
	Low       float64 `json:"3"`
	Close     float64 `json:"4"`
	Change    float64 // Thêm trường Change vào cấu trúc
}

var instance *gorm.DB

func main() {
	dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	instance = db
	log.Println("Connected to the database")

	// Gọi API
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/bitcoin/ohlc?vs_currency=usd&days=1&precision=1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Đọc dữ liệu trả về
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Phân tích dữ liệu JSON
	var ohlcData [][]interface{}
	err = json.Unmarshal(body, &ohlcData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Lưu dữ liệu vào cơ sở dữ liệu
	for i := range ohlcData {
		if i == 0 {
			continue // Bỏ qua mảng đầu tiên vì đây là tiêu đề
		}
		ohlc := OHLCData{
			Symbol:    "bitcoin",
			Timestamp: int64(ohlcData[i][0].(float64)),
			Open:      ohlcData[i][1].(float64),
			High:      ohlcData[i][2].(float64),
			Low:       ohlcData[i][3].(float64),
			Close:     ohlcData[i][4].(float64),
		}

		// Tính toán change
		if i > 0 {
			ohlc.Change = (ohlc.Close - ohlcData[i-1][4].(float64)) / ohlcData[i-1][4].(float64) * 100
		} else {
			ohlc.Change = 0
		}

		// Chuyển đổi timestamp sang định dạng thời gian bình thường
		ohlcTime := time.Unix(ohlc.Timestamp/1000, 0)
		fmt.Printf("Timestamp: %s, Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f, Change: %.2f%%\n",
			ohlcTime.Format("02 January 2006 15:04:05"), ohlc.Open, ohlc.High, ohlc.Low, ohlc.Close, ohlc.Change)

		// Lưu dữ liệu vào cơ sở dữ liệu
		if err := saveOHLCData(ohlc); err != nil {
			fmt.Println("Error saving OHLC data:", err)
		}
	}

	fmt.Println("Data saved successfully.")
}

func saveOHLCData(data OHLCData) error {
	// Thêm dữ liệu vào cơ sở dữ liệu
	if err := instance.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
