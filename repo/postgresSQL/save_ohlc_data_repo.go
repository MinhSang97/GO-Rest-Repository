package postgresSQL

import (
	"app/model"
	"gorm.io/gorm"
	"log"
)

type ohldSaveDataRepository struct {
	db *gorm.DB
}

func (s *ohldSaveDataRepository) SaveOhlcData(data []model.OHLCDataSaveData) error {

	ohldData := data

	for _, d := range ohldData {
		var count int64
		s.db.Model(&model.OHLCDataSaveData{}).Where("id = ? AND timestamp = ?", d.ID, d.Timestamp).Count(&count)
		if count == 0 {
			result := s.db.Create(&d)
			if result.Error != nil {
				return result.Error
			}
		} else {
			log.Printf("Record with id '%s' and timestamp '%d' already exists. Skipping...", d.ID, d.Timestamp)
		}
	}

	return nil
}

var instances ohldSaveDataRepository

// Corrected return type to match the interface
func SaveOhldDataRepository(db *gorm.DB) *ohldSaveDataRepository {
	if instances.db == nil {
		instances.db = db
	}
	return &instances
}
