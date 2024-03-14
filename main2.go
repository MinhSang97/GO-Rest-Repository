package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

var instances *gorm.DB

func main() {
	dsn := "host=localhost user=admin password=123456 dbname=golang port=5432 sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	instances = db
	log.Println("Connected to the database")

	// Migrate the schema
	db.AutoMigrate(&Coins{})

	url := "https://api.coingecko.com/api/v3/coins/list?include_platform=true"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var coins []Coins
	err = json.Unmarshal(body, &coins)
	if err != nil {
		log.Fatal(err)
	}

	// Lưu dữ liệu vào cơ sở dữ liệu
	for _, coin := range coins {
		db.Create(&coin)
		fmt.Printf("ID: %s, Symbol: %s, Name: %s\n", coin.ID, coin.Symbol, coin.Name)
	}
}
