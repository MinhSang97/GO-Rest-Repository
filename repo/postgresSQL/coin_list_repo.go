package postgresSQL

import (
	"app/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type coinDataRepository struct {
	db *gorm.DB
}

func (s coinDataRepository) InsertCoin(ctx context.Context, coin *model.Coins) error {

	var coins []model.Coins
	var existingCoin model.Coins
	// Find the coin based on id, symbol, and name
	for _, coin := range coins {

		result := s.db.Where("id = ?", coin.ID).First(&existingCoin)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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

	return nil
}

var instancecoin coinDataRepository

func CoinDataRepository(db *gorm.DB) coinDataRepository {
	if instancecoin.db == nil {
		instancecoin.db = db
	}
	return instancecoin
}
