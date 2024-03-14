package usecases

import (
	"app/model"
)

type InsertCoinUsecase interface {
	InsertCoin(coins []model.Coins) error
}
