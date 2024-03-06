package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/postgresSQL"
	"context"
	"time"
)

type historiesUseCase struct {
	historiesRepo repo.HistoriesRepo
}

func NewHistoriesUseCase() *historiesUseCase {
	db := dbutil.ConnectDB()
	requestGetHistoriesRepo := postgresSQL.OhlcDataRepository(db)
	return &historiesUseCase{
		historiesRepo: requestGetHistoriesRepo,
	}
}

func (uc *historiesUseCase) GetHistories(ctx context.Context, startTime time.Time, endTime time.Time, period, symbol string) ([]model.OHLCData, error) {
	return uc.historiesRepo.GetHistories(ctx, startTime.Unix(), endTime.Unix(), period, symbol)
}
