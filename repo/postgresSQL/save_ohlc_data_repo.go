//package postgresSQL
//
//import (
//	"app/model"
//	"app/repo"
//	"gorm.io/gorm"
//)
//
//type ohldSaveDataRepository struct {
//	db *gorm.DB
//}
//
//func (s ohldDataRepository) SaveOhlcData() ([]model.OHLCDataSaveData, error) {
//	var ohldData []model.OHLCDataSaveData
//
//	return ohldData, nil
//}
//
//var instances ohldSaveDataRepository
//
//// Sửa đổi chữ ký trả về của hàm OhlcDataRepository để phù hợp với repo.HistoriesRepo
//func SaveOhldDataRepository(db *gorm.DB) repo.OhlcSaveRepo {
//	if instances.db == nil {
//		instances.db = db
//	}
//	return instances
//}
