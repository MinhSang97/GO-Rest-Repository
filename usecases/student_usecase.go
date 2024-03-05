package usecases

import (
	"app/dbutil"
	"app/model"
	"app/repo"
	"app/repo/postgresSQL"
	"context"
)

type studentUseCase struct {
	studentRepo repo.StudentRepo
}

func NewStudentUseCase() StudentUsecase {
	db := dbutil.ConnectDB()
	studentRepo := postgresSQL.NewStudentRepository(db)
	return &studentUseCase{
		studentRepo: studentRepo,
	}
}

// Tested
func (uc *studentUseCase) GetOneByID(ctx context.Context, StartTime, EndTime, Period, Symbol string) (model.Student, error) {
	return uc.studentRepo.GetOneByID(ctx, StartTime, EndTime, Period, Symbol)
}
