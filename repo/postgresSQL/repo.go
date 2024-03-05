package postgresSQL

import (
	"app/model"
	"app/redis"
	"app/repo"
	"context"
	"encoding/json"
	"fmt"

	"log"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func (s studentRepository) GetOneByID(ctx context.Context, StartTime, EndTime, Period, Symbol string) (model.Student, error) {
	var student model.Student
	RedisClient := redis.ConnectRedis()

	// Đọc sinh viên từ Redis (nếu có)
	cachedStudentJSON, err := RedisClient.Get(ctx, fmt.Sprintf("student:%s", StartTime)).Result()
	if err == nil {
		var cachedStudent model.Student
		err := json.Unmarshal([]byte(cachedStudentJSON), &cachedStudent)
		if err != nil {
			log.Println("Failed to unmarshal student from Redis:", err)
			return student, fmt.Errorf("Failed to unmarshal student from Redis: %w", err)
		}
		log.Println("Student fetched from Redis")
		// Handle the response here, e.g., log or return a success message
		return cachedStudent, nil
	}

	// Nếu không tìm thấy trong Redis, đọc từ cơ sở dữ liệu MySQL
	result := s.db.First(&student, StartTime)
	if result.Error != nil {
		log.Println("Failed to fetch student from MySQL:", result.Error)
		return student, fmt.Errorf("Failed to fetch student from MySQL: %w", result.Error)
	}

	if student.ID == 0 {
		log.Println("Student not found in MySQL")
		// Handle the response here, e.g., log or return a not found message
		return student, fmt.Errorf("Student not found in MySQL")
	}

	log.Println("Student query from MySQL")

	// Cache thông tin sinh viên vào Redis
	jsonStudent, err := json.Marshal(student)
	if err != nil {
		log.Println("Failed to marshal student:", err)
		return student, fmt.Errorf("Failed to marshal student: %w", err)
	}

	key := fmt.Sprintf("student:%s", StartTime)
	err = redis.RedisClient.Set(ctx, key, jsonStudent, 0).Err()
	if err != nil {
		log.Println("Failed to cache student in Redis:", err)
		// Handle the response here, e.g., log or return an error message
	}

	// Handle the response here, e.g., log or return the student
	return student, nil
}

var instance studentRepository

func NewStudentRepository(db *gorm.DB) repo.StudentRepo {
	if instance.db == nil {
		instance.db = db

	}
	return instance
}
