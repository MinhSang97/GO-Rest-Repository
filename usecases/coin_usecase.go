package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo/postgresSQL"
	"context"
)

type coinUseCase struct {
	coinRepo InsertCoinUsecase
}

func NewCoinUseCase() *coinUseCase {
	db := dbutil.ConnectDB()
	coinRepo := postgresSQL.CoinDataRepository(db)
	return &coinUseCase{
		coinRepo: coinRepo,
	}
}

func (uc *coinUseCase) InsertCoin(ctx context.Context, coins *model.Coins) error {
	return uc.coinRepo.InsertCoin(ctx, coins)
}
