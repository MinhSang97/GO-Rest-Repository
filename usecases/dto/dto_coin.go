package dto

import (
	"app/payload"
)

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Coin) ToPayload() *payload.Coin {
	coin := &payload.Coin{
		ID:     c.ID,
		Symbol: c.Symbol,
		Name:   c.Name,
	}

	return coin
}
