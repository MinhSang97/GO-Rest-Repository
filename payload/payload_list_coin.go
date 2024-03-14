package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Coins) ToModel() model.Coins {
	coins := model.Coins{
		ID:     c.ID,
		Symbol: c.Symbol,
		Name:   c.Name,
	}
	return coins
}

func (c *Coins) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
