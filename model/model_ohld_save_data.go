package model

import (
	"encoding/json"
	"log"
)

type OHLCDataSaveData struct {
	ID        string  `json:"-"`
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Change    float64 `json:"change"`
}

func (c *OHLCDataSaveData) TableName() string {
	return "GetHistories"
}

func (c *OHLCDataSaveData) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}

func (c *OHLCDataSaveData) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
