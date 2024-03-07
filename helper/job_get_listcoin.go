package helper

import (
	"app/model"
	"app/payload"
	"app/usecases"
	"context"
	"encoding/json"
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

	var coins []model.Coins // Change to slice of model.Coins
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatal(err)
	}
	var data = payload.Coins{}
	coin := data.ToModel()
	uc := usecases.NewCoinUseCase()

	// Trích xuất thông tin từ mỗi đồng coin và chuyển vào hàm InsertCoin
	err = uc.InsertCoin(context.TODO(), coin)

	if err != nil {
		log.Println("Error inserting coin:", err)

	}

}
