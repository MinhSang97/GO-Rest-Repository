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
	// Call API
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	var ex []model.Coins
	var coins = ex
	if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
		log.Fatal(err)
	}

	uc := usecases.NewCoinUseCase()

	// Iterate over each coin and insert it

	var data = payload.Coins{} // Assuming this is how you convert model.Coins to payload.Coins
	coinPayload := data.ToModel()

	// Insert coin
	err = uc.InsertCoin(context.TODO(), coinPayload)
	if err != nil {
		log.Println("Error inserting coin:", err)
	}

}
