package postgresSQL

import (
	"app/model"
	"app/repo"
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type ohldDataRepository struct {
	db *gorm.DB
}

func (s ohldDataRepository) GetHistories(ctx context.Context, startDate int64, endDate int64, period, symbol string) ([]model.OHLCData, error) {
	var ohldData []model.OHLCData

	switch period {
	case "30M", "1H", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "11H", "12H", "13H", "14H", "15H", "16H", "17H", "18H", "19H", "20H", "21H", "22H", "23H", "24H":
		var ohlcData []model.OHLCData
		err := s.db.Raw("SELECT * FROM ohlc_data WHERE timestamp = ? OR timestamp = ?", startDate, endDate).Error
		if err != nil {
			log.Println("Failed to execute SQL query:", err)
			return ohlcData, fmt.Errorf("Failed to execute SQL query: %w", err)

		} else {
			rows, err := s.db.Raw("SELECT o.timestamp, o.high, o.low, o.open, o.close, o.change FROM coins c RIGHT JOIN ohlc_data o ON c.id = o.id WHERE c.symbol = ? AND o.timestamp >= ? AND o.timestamp <= ?", symbol, startDate, endDate).Rows()
			if err != nil {
				log.Println("Failed to execute SQL query:", err)
				return ohldData, fmt.Errorf("Failed to execute SQL query: %w", err)
			}
			defer rows.Close()

			for rows.Next() {
				var timestamp int64
				var high, low, open, close, change float64
				rows.Scan(&timestamp, &high, &low, &open, &close, &change)
				ohldData = append(ohldData, model.OHLCData{Timestamp: timestamp, High: high, Low: low, Open: open, Close: close, Change: change})
			}

			if len(ohldData) == 0 {
				log.Println("Data not found in PostgreSQL")
				return ohldData, fmt.Errorf("Data not found in PostgreSQL")
			}

			var highestPrice, lowestPrice, firstPrice, lastPrice float64
			var startTime int64

			for i, data := range ohldData {
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

			var change float64
			if len(ohldData) > 1 {
				change = (lastPrice - ohldData[len(ohldData)-2].Close) / ohldData[len(ohldData)-2].Close * 100
			}

			data := model.OHLCData{
				Timestamp: startTime,
				High:      highestPrice,
				Low:       lowestPrice,
				Open:      firstPrice,
				Close:     lastPrice,
				Change:    change,
			}

			log.Println("Data query from PostgreSQL")

			datax := []model.OHLCData{}

			ohldData = append(datax, data)
		}
		return ohldData, nil

	case "2D", "3D", "4D", "5D", "6D", "7D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	default:
		return ohldData, fmt.Errorf("chưa hỗ trợ")

	}

	log.Println("Data query from PostgreSQL")

	return ohldData, nil
}

var instance ohldDataRepository

// Sửa đổi chữ ký trả về của hàm OhlcDataRepository để phù hợp với repo.HistoriesRepo
func OhlcDataRepository(db *gorm.DB) repo.HistoriesRepo {
	if instance.db == nil {
		instance.db = db
	}
	return instance
}
