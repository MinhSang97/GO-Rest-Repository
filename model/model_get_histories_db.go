package model

import (
	"encoding/json"
	"log"
)

type GetHistories struct {
	StartTime string `json:"start_date" binding:"required"`
	EndTime   string `json:"end_date" binding:"required"`
	Period    string `json:"period" binding:"required"`
	Symbol    string `json:"symbol" binding:"required"`
}

func (c *GetHistories) TableName() string {
	return "GetHistories"
}

func (c *GetHistories) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *GetHistories) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
