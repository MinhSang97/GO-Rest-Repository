package dto

import (
	"app/payload"
)

type Coins struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Coins) ToPayload() *payload.Coins {
	coins := &payload.Coins{
		ID:     c.ID,
		Symbol: c.Symbol,
		Name:   c.Name,
	}

	return coins
}
