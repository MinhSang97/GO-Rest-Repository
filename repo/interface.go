package repo

import (
	"app/model"
	"context"
)

type StudentRepo interface {
	GetOneByID(ctx context.Context, StartTime, EndTime, Period, Symbol string) (model.Student, error)
}
