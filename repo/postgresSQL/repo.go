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

// Chuyển startDate và endDate thành startTime và endTime, và chỉ trả về một bản ghi OHLCData
//func (s ohldDataRepository) GetHistories(ctx context.Context, startDate int64, endDate int64, period, symbol string) (model.OHLCData, error) {
//	var ohldData model.OHLCData
//
//	// Nếu không tìm thấy trong Redis, thực hiện truy vấn vào PostgreSQL
//	// Sử dụng câu lệnh SQL tương ứng với yêu cầu
//	rows, err := s.db.Raw("SELECT o.timestamp, o.high, o.low, o.open, o.close, o.change FROM coins c RIGHT JOIN ohlc_data o ON c.id = o.id WHERE c.symbol = ? AND o.timestamp >= ? AND o.timestamp <= ?", symbol, startDate, endDate).Rows()
//	if err != nil {
//		log.Println("Failed to execute SQL query:", err)
//		return ohldData, fmt.Errorf("Failed to execute SQL query: %w", err)
//	}
//	defer rows.Close()
//
//	// Xử lý kết quả từ PostgreSQL
//	for rows.Next() {
//		var timestamp int64
//		var high, low, open, close, change float64
//		rows.Scan(&timestamp, &high, &low, &open, &close, &change)
//		ohldData = model.OHLCData{Timestamp: timestamp, High: high, Low: low, Open: open, Close: close, Change: change}
//	}
//
//	if ohldData.Open == 0 {
//		log.Println("Data not found in PostgreSQL")
//		return ohldData, fmt.Errorf("Data not found in PostgreSQL")
//	}
//
//	log.Println("Data query from PostgreSQL")
//
//	return ohldData, nil
//}

func (s ohldDataRepository) GetHistories(ctx context.Context, startDate int64, endDate int64, period, symbol string) ([]model.OHLCData, error) {
	var ohldData []model.OHLCData

	// Nếu không tìm thấy trong Redis, thực hiện truy vấn vào PostgreSQL
	// Sử dụng câu lệnh SQL tương ứng với yêu cầu
	rows, err := s.db.Raw("SELECT o.timestamp, o.high, o.low, o.open, o.close, o.change FROM coins c RIGHT JOIN ohlc_data o ON c.id = o.id WHERE c.symbol = ? AND o.timestamp >= ? AND o.timestamp <= ?", symbol, startDate, endDate).Rows()
	if err != nil {
		log.Println("Failed to execute SQL query:", err)
		return ohldData, fmt.Errorf("Failed to execute SQL query: %w", err)
	}
	defer rows.Close()

	// Xử lý kết quả từ PostgreSQL
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
