package dbutil

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"

		var err error
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		instance = db
		log.Println("Connected to the database")
	})

	return instance
}
