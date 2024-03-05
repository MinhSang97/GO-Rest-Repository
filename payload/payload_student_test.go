package payload

import (
	"testing"
	"time"
)

func TestAddStudentRequest_ToModel(t *testing.T) {
	// Tạo một đối tượng Student để kiểm thử
	student := AddStudentRequest{
		//ID:           1,
		FirstName:    "Nguyễn",
		LastName:     "Minh Sang",
		Age:          27,
		Grade:        9.5,
		ClassName:    "hhm",
		EntranceDate: time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	toModel := student.ToModel()

	if student.FirstName != toModel.FirstName {
		t.Errorf("Expected FirstName %s, but got %s", student.FirstName, student.FirstName)
	}

	// Rest of your testing code...
}

func TestAddStudentRequest_FromJson(t *testing.T) {
	jsonData := `{"first_name":"Minh Sang","last_name":"Nguyễn","age":27,"grade":9.8,"class_name":"hhm","entrance_date":"2022-01-08T12:00:00Z","created_at":"2022-01-08T12:00:00Z","updated_at":"2022-01-08T12:00:00Z"}`
	student := AddStudentRequest{}

	student.FromJson(jsonData)

	if student.FirstName != "Minh Sang" || student.LastName != "Nguyễn" || student.Age != 27 || student.Grade != 9.8 || student.ClassName != "hhm" {
		t.Errorf("Incorrect values after parsing JSON")

	}
}
