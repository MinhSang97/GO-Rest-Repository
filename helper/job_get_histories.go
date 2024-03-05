package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type OHLCData struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Change    float64
}

func main() {
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

	// Tạo slice để lưu trữ dữ liệu OHLC
	var ohlcList []OHLCData

	// Tính toán change và lưu dữ liệu vào slice
	for i := range ohlcData {
		if i == 0 {
			continue // Bỏ qua mảng đầu tiên vì đây là tiêu đề
		}
		ohlc := OHLCData{
			Timestamp: int64(ohlcData[i][0].(float64)),
			Open:      ohlcData[i][1].(float64),
			High:      ohlcData[i][2].(float64),
			Low:       ohlcData[i][3].(float64),
			Close:     ohlcData[i][4].(float64),
		}

		// Tính toán change
		if len(ohlcList) > 0 {
			ohlc.Change = (ohlc.Close - ohlcList[len(ohlcList)-1].Close) / ohlcList[len(ohlcList)-1].Close * 100
		} else {
			ohlc.Change = 0
		}

		ohlcList = append(ohlcList, ohlc)
	}

	// In dữ liệu đã phân tích
	for _, data := range ohlcList {
		// Chuyển đổi timestamp sang thời gian bình thường
		t := time.Unix(data.Timestamp/1000, 0).UTC()
		fmt.Printf("Timestamp: %s, Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f, Change: %.2f%%\n", t.Format("02 January 2006 15:04:05"), data.Open, data.High, data.Low, data.Close, data.Change)
	}
}
