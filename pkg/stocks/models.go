package stocks

type TickerData struct {
	Metadata   TimeseriesMetadata        `json:"Meta Data"`
	Timeseries map[string]TimeseriesData `json:"Time Series (Daily)"`
	OutputData OutputData                `json:"Output Data"`
}

type OutputData struct {
	NDays    int    `json:"N Days"`
	CloseAvg string `json:"Close Average"`
}

type TimeseriesMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type TimeseriesData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}
