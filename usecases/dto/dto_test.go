// student_test.go
package dto

import (
	"testing"
	"time"
)

func TestToPayload(t *testing.T) {
	// Tạo một đối tượng Student để kiểm thử
	student := Student{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Age:          20,
		Grade:        8.5,
		ClassName:    "Math",
		EntranceDate: time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Gọi hàm ToPayload để lấy đối tượng payload
	payload := student.ToPayload()

	// Kiểm tra xem payload có được tạo đúng không
	if payload.FirstName != student.FirstName {
		t.Errorf("Expected FirstName %s, but got %s", student.FirstName, payload.FirstName)
	}

	if payload.LastName != student.LastName {
		t.Errorf("Expected LastName %s, but got %s", student.LastName, payload.LastName)
	}

	if payload.Age != student.Age {
		t.Errorf("Expected Age %d, but got %d", student.Age, payload.Age)
	}

	if payload.Grade != student.Grade {
		t.Errorf("Expected Grade %f, but got %f", student.Grade, payload.Grade)
	}

	if payload.ClassName != student.ClassName {
		t.Errorf("Expected ClassName %s, but got %s", student.ClassName, payload.ClassName)
	}

	if payload.EntranceDate != student.EntranceDate {
		t.Errorf("Expected EntranceDate %s, but got %s", student.EntranceDate, payload.EntranceDate)
	}

}
