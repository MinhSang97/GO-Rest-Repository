package payload

import (
	"app/model"
	"encoding/json"
	"log"
)

//type RequestGetHistories struct {
//	StartTime string `json:"start_time"`
//	EndTime   string `json:"end_time"`
//	Period    string `json:"period"`
//	Symbol    string `json:"symbol"`
//}

type RequestGetHistories struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Period    string `json:"period"`
	Symbol    string `json:"symbol"`
}

func (c *RequestGetHistories) ToModel() *model.RequestGetHistories {
	requestGetHistories := &model.RequestGetHistories{
		StartDate: c.StartDate,
		EndDate:   c.EndDate,
		Period:    c.Period,
		Symbol:    c.Symbol,
	}
	return requestGetHistories
}

func (c *RequestGetHistories) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Error  string      `json:"error,omitempty"`
}
