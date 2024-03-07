package main

import (
	"app/model"
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// Kết nối đến cơ sở dữ liệu PostgreSQL
	dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(model.Coins{})
	// Gọi API
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var coins []model.Coins
	// Đọc dữ liệu từ API
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatal(err)
	}
	var existingCoin model.Coins
	// Lưu dữ liệu vào cơ sở dữ liệu
	for _, coin := range coins {

		result := db.Where("id = ?", coin.ID).First(&existingCoin)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&coin).Error; err != nil {
					log.Println("Failed to insert coin:", err)
				}
			} else {
				log.Println("Failed to query coin:", result.Error)
			}
		} else {
			// Cập nhật dữ liệu nếu coin đã tồn tại trong DB
			existingCoin.Symbol = coin.Symbol
			existingCoin.Name = coin.Name
			if err := db.Save(&existingCoin).Error; err != nil {
				log.Println("Failed to update coin:", err)
			}
		}
	}

	log.Println("Data inserted/updated successfully")
}
