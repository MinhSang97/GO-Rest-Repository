package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo/postgresSQL"
)

type saveOhlcDataUseCase struct {
	ohlcSaveRepo SaveOhlcData
}

func SaveOhlcDataUseCase() *saveOhlcDataUseCase {
	db := dbutil.ConnectDB()
	ohlcSaveRepo := postgresSQL.SaveOhldDataRepository(db)
	return &saveOhlcDataUseCase{
		ohlcSaveRepo: ohlcSaveRepo,
	}
}

func (uc *saveOhlcDataUseCase) SaveOhlcData(data []model.OHLCDataSaveData) error {
	return uc.ohlcSaveRepo.SaveOhlcData(data)
}
