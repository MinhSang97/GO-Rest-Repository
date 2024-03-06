package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Coin) ToModel() *model.Coin {
	coin := &model.Coin{
		ID:     c.ID,
		Symbol: c.Symbol,
		Name:   c.Name,
	}
	return coin
}

func (c *Coin) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
