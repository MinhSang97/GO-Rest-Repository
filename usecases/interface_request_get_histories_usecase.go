package usecases

import (
	"app/model"
	"context"
)

type GetHistoriesUsecase interface {
	GetHistories(ctx context.Context, StartTime, EndTime, Period, Symbol string) (model.RequestGetHistories, error)
}
