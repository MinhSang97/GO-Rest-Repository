package repo

import (
	"app/model"
	"context"
)

type HistoriesRepo interface {
	//GetHistories(ctx context.Context, startTime int64, endTime int64, period, symbol string) (model.OHLCData, error)
	GetHistories(ctx context.Context, startTime int64, endTime int64, period string, symbol string) ([]model.OHLCData, error)
}

type CoinRepo interface {
	Coin(ctx context.Context, id, symbol, name string) (model.Coin, error)
}
