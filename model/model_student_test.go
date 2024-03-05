package model

import (
	"testing"
	"time"
)

func TestStudent_TableName(t *testing.T) {
	student := Student{}
	expected := "student"
	result := student.TableName()

	if result != expected {
		t.Errorf("Expected table name %s, but got %s", expected, result)
	}
}

func TestStudent_ToJson(t *testing.T) {
	student := Student{
		ID:           1,
		FirstName:    "Minh Sang",
		LastName:     "Nguyễn",
		Age:          27,
		Grade:        9.8,
		ClassName:    "hhm",
		EntranceDate: time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	expectedJSON := `{"id":1,"first_name":"Minh Sang","last_name":"Nguyễn","age":27,"grade":9.8,"class_name":"hhm","entrance_date":"` +
		student.EntranceDate.Format(time.RFC3339Nano) + `","created_at":"` +
		student.CreatedAt.Format(time.RFC3339Nano) + `","updated_at":"` +
		student.UpdatedAt.Format(time.RFC3339Nano) + `"}`

	resultJSON := student.ToJson()

	if resultJSON != expectedJSON {
		t.Errorf("Expected JSON %s, but got %s", expectedJSON, resultJSON)
	}
}

func TestStudent_FromJson(t *testing.T) {
	jsonData := `{"id":1,"first_name":"Minh Sang","last_name":"Nguyễn","age":27,"grade":9.8,"class_name":"hhm","entrance_date":"2022-01-08T12:00:00Z","created_at":"2022-01-08T12:00:00Z","updated_at":"2022-01-08T12:00:00Z"}`
	student := Student{}

	student.FromJson(jsonData)

	if student.ID != 1 || student.FirstName != "Minh Sang" || student.LastName != "Nguyễn" || student.Age != 27 || student.Grade != 9.8 || student.ClassName != "hhm" {
		t.Errorf("Incorrect values after parsing JSON")

	}
}
