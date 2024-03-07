package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/postgresSQL"
	"context"
)

type coinUseCase struct {
	coinRepo repo.CoinRepo
}

func NewCoinUseCase() *coinUseCase {
	db := dbutil.ConnectDB()
	coinRepo := postgresSQL.CoinDataRepository(db)
	return &coinUseCase{
		coinRepo: coinRepo,
	}
}

func (uc *coinUseCase) InsertCoin(ctx context.Context, id string, symbol string, name string) ([]model.Coin, error) {
	return uc.InsertCoin(ctx, id, "", "")
}
