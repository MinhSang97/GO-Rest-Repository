package helper

import (
	"app/model"
	"app/usecases"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ListCoin() {
	//// Call API
	//response, err := http.Get("https://api.coingecko.com/api/v3/coins/list?include_platform=true")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer response.Body.Close()
	//var ex []model.Coins
	//var coins = ex
	//if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
	//	log.Fatal(err)
	//}
	//
	//uc := usecases.NewCoinUseCase()
	//
	//// Iterate over each coin and insert it
	//
	//var data = payload.Coins{} // Assuming this is how you convert model.Coins to payload.Coins
	//coinPayload := data.ToModel()
	//
	//// Insert coin
	//err = uc.InsertCoin(context.TODO(), coinPayload)
	//if err != nil {
	//	log.Println("Error inserting coin:", err)
	//}

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

	var coins []model.Coins
	err = json.Unmarshal(body, &coins)
	if err != nil {
		log.Fatal(err)
	}
	uc := usecases.NewCoinUseCase()

	// Iterate over each coin and insert it

	// Insert coins
	err = uc.InsertCoin(coins)
	if err != nil {
		log.Println("Error inserting coins:", err)
	}

}
