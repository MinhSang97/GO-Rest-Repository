package helper

import (
	"app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ListCoin() {
	// Gọi API
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var coins []model.Coin // Đọc dữ liệu từ API
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatal(err)
	}

	// In ra dữ liệu JSON
	for _, coin := range coins {
		coinJSON, err := json.Marshal(coin)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			continue
		}
		fmt.Println(string(coinJSON))
	}
}
