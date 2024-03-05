package usecases

import (
	"app/model"
	"context"
)

type StudentUsecase interface {
	GetOneByID(ctx context.Context, StartTime, EndTime, Period, Symbol string) (model.Student, error)
}
