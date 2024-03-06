package model

import (
	"encoding/json"
	"log"
)

type RequestGetHistories struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Period    string `json:"period" binding:"required"`
	Symbol    string `json:"symbol" binding:"required"`
}

func (c *RequestGetHistories) TableName() string {
	return "RequestGetHistories"
}

func (c *RequestGetHistories) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *RequestGetHistories) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
