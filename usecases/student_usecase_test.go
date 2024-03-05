// student_usecase_test.go
package usecases

import (
	"app/model"
	mockStudentRepo "app/repo/mocks"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStudentUseCase_GetStudentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	expectedStudent := model.Student{ID: 53}
	mockRepo.EXPECT().GetOneByID(gomock.Any(), 53).Return(expectedStudent, nil)

	// Thực hiện unit test
	student, err := useCase.GetOneByID(context.Background(), 53)

	assert.Nil(t, err)
	assert.Equal(t, expectedStudent, student)
}

func TestStudentUseCase_GetAllStudents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	expectedStudents := []model.Student{
		{ID: 13},
		{ID: 14},
	}
	mockRepo.EXPECT().GetAll(gomock.Any()).Return(expectedStudents, nil)

	// Thực hiện unit test
	students, err := useCase.GetAll(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, expectedStudents, students)
}

func TestStudentUseCase_CreateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	mockRepo.EXPECT().CreateStudent(gomock.Any(), gomock.Any()).Return(nil)

	// Thực hiện unit test
	student := &model.Student{ID: 15}
	err := useCase.CreateStudent(context.Background(), student)

	assert.Nil(t, err)
}

func TestStudentUseCase_UpdateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	mockRepo.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	// Thực hiện unit test
	student := &model.Student{ID: 15}
	err := useCase.UpdateOne(context.Background(), 15, student)

	assert.Nil(t, err)
}

func TestStudentUseCase_DeleteStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	//student := model.Student{ID: 15}
	mockRepo.EXPECT().DeleteOne(gomock.Any(), gomock.Any()).Return(nil)

	// Thực hiện unit test
	err := useCase.DeleteOne(context.Background(), 15)

	//// Thiết lập behavior cho mockRepo
	//expectedStudent := model.Student{ID: 53}
	//mockRepo.EXPECT().GetOneByID(gomock.Any(), 53).Return(expectedStudent, nil)
	//
	//// Thực hiện unit test
	//student, err := useCase.GetOneByID(context.Background(), 53)

	assert.Nil(t, err)
}

func TestStudentUseCase_SearchStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockStudentRepo.NewMockStudentRepo(ctrl)
	useCase := studentUseCase{
		studentRepo: mockRepo,
	}

	// Thiết lập behavior cho mockRepo
	expectedResult := []model.Student{} // Modify the expected result accordingly
	mockRepo.EXPECT().Search(gomock.Any(), gomock.Any()).Return(expectedResult, nil)

	// Thực hiện unit test
	result, err := useCase.Search(context.Background(), "test")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result) // Add an assertion to check the result
}
