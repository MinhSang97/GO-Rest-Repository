package postgresSQL

import (
	"app/model"
	"app/repo"
	"context"
	"gorm.io/gorm"
	"log"
)

type coinDataRepository struct {
	db *gorm.DB
}

func (s coinDataRepository) Coin(ctx context.Context, id, symbol, name string) ([]model.Coin, error) {
	var coinData []model.Coin

	// Auto Migrate Schema
	s.db.AutoMigrate(&model.Coin{})

	// Lưu dữ liệu vào cơ sở dữ liệu
	for _, coin := range coinData {
		var existingCoin model.Coin
		result := s.db.Where("id = ?", coin.ID).First(&existingCoin)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := s.db.Create(&coin).Error; err != nil {
					log.Println("Failed to insert coin:", err)
				}
			} else {
				log.Println("Failed to query coin:", result.Error)
			}
		} else {
			// Cập nhật dữ liệu nếu coin đã tồn tại trong DB
			existingCoin.Symbol = coin.Symbol
			existingCoin.Name = coin.Name
			if err := s.db.Save(&existingCoin).Error; err != nil {
				log.Println("Failed to update coin:", err)
			}
		}
	}

	log.Println("Data inserted/updated successfully")

	return coinData, nil
}

//func (s coinDataRepository) InsertCoin(ctx context.Context, id, symbol, name string) ([]model.Coin, error) {
//	var coinData []model.Coin
//
//	// Auto Migrate Schema
//	s.db.AutoMigrate(&model.Coin{})
//
//	// Lưu dữ liệu vào cơ sở dữ liệu
//	for _, coin := range coinData {
//		var existingCoin model.Coin
//		result := s.db.Where("id = ?", coin.ID).First(&existingCoin)
//		if result.Error != nil {
//			if result.Error == gorm.ErrRecordNotFound {
//				if err := s.db.Create(&coin).Error; err != nil {
//					log.Println("Failed to insert coin:", err)
//				}
//			} else {
//				log.Println("Failed to query coin:", result.Error)
//			}
//		} else {
//			// Cập nhật dữ liệu nếu coin đã tồn tại trong DB
//			existingCoin.Symbol = coin.Symbol
//			existingCoin.Name = coin.Name
//			if err := s.db.Save(&existingCoin).Error; err != nil {
//				log.Println("Failed to update coin:", err)
//			}
//		}
//	}
//
//	log.Println("Data inserted/updated successfully")
//
//	return coinData, nil
//}

var instancecoin coinDataRepository

func CoinDataRepository(db *gorm.DB) repo.CoinRepo {
	if instancecoin.db == nil {
		instancecoin.db = db
	}
	return instancecoin
}
