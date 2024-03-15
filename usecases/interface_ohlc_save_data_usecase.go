package usecases

import (
	"app/model"
)

type SaveOhlcData interface {
	SaveOhlcData(coins []model.OHLCDataSaveData) error
}
