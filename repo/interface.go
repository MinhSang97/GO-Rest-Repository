package repo

import (
	"app/model"
	"context"
)

type HistoriesRepo interface {
	GetHistories(ctx context.Context, startTime int64, endTime int64, period string, symbol string) ([]model.OHLCData, error)
}

type CoinRepo interface {
	Coin(coins *model.Coins) error
}

//type OhlcSaveRepo interface {
//	SaveOhlcData(saveOhlcData *model.OHLCDataSaveData) ([]model.OHLCDataSaveData, error)
//}

//saveOhlcData *model.OHLCDataSaveData
