package dto

import (
	"app/payload"
)

type RequestGetHistories struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Period    string `json:"period"`
	Symbol    string `json:"symbol"`
}

func (c *RequestGetHistories) ToPayload() *payload.RequestGetHistories {
	requestGetHistories := &payload.RequestGetHistories{
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		Period:    c.Period,
		Symbol:    c.Symbol,
	}

	return requestGetHistories
}
