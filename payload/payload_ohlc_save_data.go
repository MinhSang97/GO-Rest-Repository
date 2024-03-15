package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

type OHLCDataSaveData struct {
	ID        string  `json:"id"`
	Timestamp int64   `json:"timestamp"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	Change    float64 `json:"change"`
}

func (c *OHLCDataSaveData) ToModel() *model.OHLCDataSaveData {
	ohlcDataSaveData := &model.OHLCDataSaveData{
		Timestamp: c.Timestamp,
		High:      c.High,
		Low:       c.Low,
		Open:      c.Open,
		Close:     c.Close,
		Change:    c.Change,
	}
	return ohlcDataSaveData
}

func (c *OHLCDataSaveData) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
