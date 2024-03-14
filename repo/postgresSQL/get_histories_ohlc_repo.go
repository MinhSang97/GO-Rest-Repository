package postgresSQL

import (
	"app/model"
	"app/redis"
	"app/repo"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type ohldDataRepository struct {
	db *gorm.DB
}

var RedisClient = redis.ConnectRedis()

func (s ohldDataRepository) GetHistories(ctx context.Context, startDate int64, endDate int64, period, symbol string) ([]model.OHLCData, error) {
	var ohldData []model.OHLCData

	switch period {
	case "1D":
		var ohlcData []model.OHLCData
		var count int
		err := s.db.Raw("SELECT * FROM ohlc_data_1day WHERE timestamp = ? OR timestamp = ?", startDate, endDate).Error
		if err != nil || count != 2 {
			log.Println("End Time not available. Need call API get data")
			return ohlcData, fmt.Errorf("End Time not available. Need call API get data")

		} else {
			cacheKey := fmt.Sprintf("ohlc_data:%s:%d:%d", symbol, startDate, endDate)

			// Kiểm tra cache trong Redis
			cachedOhldDataJSON, err := RedisClient.Get(ctx, cacheKey).Result()
			if err == nil {
				var cachedOhldData []model.OHLCData
				err := json.Unmarshal([]byte(cachedOhldDataJSON), &cachedOhldData)
				if err != nil {
					log.Println("Failed to unmarshal data from Redis:", err)
					return ohldData, fmt.Errorf("Failed to unmarshal data from Redis: %w", err)
				}
				log.Println("Data fetched from Redis")
				return cachedOhldData, nil
			}

			// Truy vấn dữ liệu từ PostgreSQL
			rows, err := s.db.Raw("SELECT o.timestamp, o.high, o.low, o.open, o.close, o.change FROM coins c RIGHT JOIN ohlc_data_1day o ON c.id = o.id WHERE c.symbol = ? AND o.timestamp >= ? AND o.timestamp <= ?", symbol, startDate, endDate).Rows()
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

			if len(ohldData) == 0 {
				log.Println("Data not found in PostgreSQL")
				return ohldData, fmt.Errorf("Data not found in PostgreSQL")
			}

			// Cache dữ liệu vào Redis
			cacheDuration := time.Hour
			ohldDataJSON, err := json.Marshal(ohldData)
			if err != nil {
				log.Println("Failed to marshal data for Redis cache:", err)
				return ohldData, fmt.Errorf("Failed to marshal data for Redis cache: %w", err)
			}
			err = RedisClient.Set(ctx, cacheKey, ohldDataJSON, cacheDuration).Err()
			if err != nil {
				log.Println("Failed to cache data in Redis:", err)
				// Không trả về lỗi ở đây vì dữ liệu đã được trả về thành công từ PostgreSQL
			}

			log.Println("Data queried from PostgreSQL and cached in Redis")

		}
		return ohldData, nil
	case "7D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "14D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "30D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "90D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "180D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "365D":
		return ohldData, fmt.Errorf("chưa hỗ trợ")
	case "MAX":
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
