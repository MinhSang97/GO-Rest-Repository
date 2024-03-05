package payload

type RequestGetHistories struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Period    string `json:"period"`
	Symbol    string `json:"symbol"`
}

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Error  string      `json:"error,omitempty"`
}
