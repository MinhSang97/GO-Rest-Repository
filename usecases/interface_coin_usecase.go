package usecases

import (
	"app/model"
	"context"
)

type InsertCoinUsecase interface {
	InsertCoin(ctx context.Context, coins *model.Coins) error
}
