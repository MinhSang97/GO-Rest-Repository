//package usecases
//
//import (
//	"app/dbutil"
//	"app/model"
//	"app/repo"
//	"app/repo/postgresSQL"
//)
//
//type saveOhlcDataUseCase struct {
//	ohlcSaveRepo repo.OhlcSaveRepo
//}
//
//func SaveOhlcDataUseCase() *saveOhlcDataUseCase {
//	db := dbutil.ConnectDB()
//	saveOhlcDataRepo := postgresSQL.SaveOhldDataRepository(db)
//	return &saveOhlcDataUseCase{
//		ohlcSaveRepo: saveOhlcDataRepo,
//	}
//}
//
//func (uc *saveOhlcDataUseCase) SaveOhlcData(*model.OHLCDataSaveData) ([]model.OHLCDataSaveData, error) {
//	return uc.ohlcSaveRepo.SaveOhlcData(*model.OHLCDataSaveData)
//}
