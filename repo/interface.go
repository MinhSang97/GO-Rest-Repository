package repo

import (
	"app/model"
	"context"
)

type HistoriesRepo interface {
	GetHistories(ctx context.Context, startTime int64, endTime int64, period string, symbol string) ([]model.OHLCData, error)
}

type CoinRepo interface {
	Coin(ctx context.Context, coins *model.Coins) error
}
