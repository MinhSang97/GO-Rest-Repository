package usecases

import (
	"app/model"
	"context"
)

type InsertCoinUsecase interface {
	GetHistories(ctx context.Context, id, symbol, name string) ([]model.Coin, error)
}
