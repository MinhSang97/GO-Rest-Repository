package model

import (
	"encoding/json"
	"log"
)

type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Coins) TableName() string {
	return "coins"
}

func (c *Coins) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *Coins) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
