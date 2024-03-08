package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type OHLCData struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Change    float64 `json:"change"`
}

func main() {
	// Chuỗi JSON đầu vào
	var jsonData = `[
        {
            "id": "",
            "timestamp": 1709787600,
            "open": 66165.8,
            "high": 66165.8,
            "low": 65866.5,
            "close": 65866.5,
            "change": -0.24972967406414856
        },
        {
            "id": "",
            "timestamp": 1709789400,
            "open": 65868.6,
            "high": 66120,
            "low": 65868.6,
            "close": 66120,
            "change": 0.38486939491243655
        },
        {
            "id": "",
            "timestamp": 1709791200,
            "open": 66049.2,
            "high": 66049.2,
            "low": 65784.1,
            "close": 65784.1,
            "change": -0.5080157289776076
        }
    ]`

	// Chuyển đổi chuỗi JSON thành slice các cấu trúc dữ liệu OHLCData
	var ohlcData []OHLCData
	err := json.Unmarshal([]byte(jsonData), &ohlcData)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// Tính toán các giá trị mới
	var highestPrice, lowestPrice, firstPrice, lastPrice float64
	var startTime int64

	// Tìm giá trị cao nhất, thấp nhất, giá mở cửa và đóng cửa
	for i, data := range ohlcData {
		if i == 0 {
			firstPrice = data.Open
			startTime = data.Timestamp
		}
		if data.High > highestPrice {
			highestPrice = data.High
		}
		if data.Low < lowestPrice || lowestPrice == 0 {
			lowestPrice = data.Low
		}
		lastPrice = data.Close
	}

	// Tính toán phần trăm thay đổi
	var change float64
	if len(ohlcData) > 1 {
		change = (lastPrice - ohlcData[len(ohlcData)-2].Close) / ohlcData[len(ohlcData)-2].Close * 100
	}

	// Tạo cấu trúc mới với các giá trị tính toán
	data := map[string]interface{}{
		"high":   highestPrice,
		"low":    lowestPrice,
		"open":   firstPrice,
		"close":  lastPrice,
		"time":   startTime,
		"change": change,
	}

	// Chuyển đổi cấu trúc mới thành chuỗi JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error converting to JSON:", err)
	}

	fmt.Println(string(dataJSON))
}
