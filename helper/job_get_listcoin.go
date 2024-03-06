package main

import (
	"encoding/json"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func main() {
	// Kết nối đến cơ sở dữ liệu PostgreSQL
	dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate Schema
	db.AutoMigrate(&Coin{})

	// Gọi API
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var coins []Coin
	// Đọc dữ liệu từ API
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatal(err)
	}

	// Lưu dữ liệu vào cơ sở dữ liệu
	for _, coin := range coins {
		if err := db.Create(&coin).Error; err != nil {
			log.Println("Failed to insert coin:", err)
		}
	}

	log.Println("Data inserted successfully")
}
