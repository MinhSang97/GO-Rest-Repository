package main

import (
	"app/payload"
	"encoding/json"
	"fmt"
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

// Hàm để tính toán khoảng thời gian dựa trên giá trị của Period
//func addTimeForPeriod(period string) string {
//	// Thêm đơn vị "h" vào chuỗi Period trước khi chuyển đổi
//
//	var num int
//	var unit string
//	fmt.Sscanf(period, "%d%s", &num, &unit)
//	fmt.Println(num)
//
//	periodWithUnit := period + "h"
//	fmt.Println(periodWithUnit)
//
//	return unit
//}

func fetchData() ([]OHLCData, error) {

	request := payload.RequestGetHistories{}
	//loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh") // Lấy múi giờ của Việt Nam
	//startDate, err := time.ParseInLocation("02-01-2006 15:04:05", request.StartDate, loc)
	//
	//endDate, err := time.ParseInLocation("02-01-2006 15:04:05", request.EndDate, loc)

	//duration := endDate.Sub(startDate)
	//days := int(duration.Hours() / 24)
	//if duration.Hours()%24 > 0 {
	//	days++ // Làm tròn lên nếu phần thập phân lớn hơn 0
	//}

	//duration := endDate.Sub(startDate)
	//days := int(duration.Hours() / 24)
	//hours := int(duration.Hours()) % 24 // Lấy phần dư của số giờ

	// Nếu số giờ lớn hơn 0, làm tròn lên số ngày
	//if hours > 0 {
	//	days++
	//}
	//day := days
	//fmt.Println("days: ", days)

	//switch request.Period {
	//case "30M", "1H", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "11H", "12H", "13H", "14H", "15H", "16H", "17H", "18H", "19H", "20H", "21H", "22H", "23H", "24H":
	//day := 1
	//case "2D", "3D", "4D", "5D", "6D", "7D":
	//
	//default:
	//
	//}
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

	var ohlcData []OHLCData
	for _, d := range data {
		// Ensure each element is an array of length 5
		if len(d) != 5 {
			return nil, fmt.Errorf("unexpected data format")
		}

		ohlc := OHLCData{
			ID:        "bitcoin",
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

func calculateChange(data *[]OHLCData) {
	for i := range *data {
		if i > 0 {
			previousClose := (*data)[i-1].Close
			(*data)[i].Change = (((*data)[i].Close - previousClose) / previousClose) * 100
		} else {
			(*data)[i].Change = 0
		}
	}
}

func saveData(data *[]OHLCData) error {
	for _, d := range *data {
		// Check if the record already exists
		var count int64
		instance.Model(&OHLCData{}).Where("id = ? AND timestamp = ?", d.ID, d.Timestamp).Count(&count)
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
