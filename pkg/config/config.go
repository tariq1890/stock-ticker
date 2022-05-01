package config

import (
	"github.com/alecthomas/kingpin"
)

const (
	defaultSymbol = "MSFT"
	defaultAPIURL = "https://www.alphavantage.co/query"
)

type Config struct {
	Symbol string
	NDays  int
	APIKey string
	APIURL string
}

var defaultConfig = &Config{
	Symbol: defaultSymbol,
	APIURL: defaultAPIURL,
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags(args []string) error {
	app := kingpin.New("stock-ticker", "StockTicker is a webservice that calls the AlphaVantage API internally to return stock data given N days and a stock symbol")
	app.DefaultEnvars()

	app.Flag("symbol", "The Stock Symbol to query (default: FORG)").Default(defaultConfig.Symbol).StringVar(&cfg.Symbol)
	app.Flag("api-host", "The API endpoint to use (default: https://www.alphavantage.co/query)").Default(defaultConfig.APIURL).StringVar(&cfg.APIURL)
	app.Flag("api-key", "The API Key to use when querying the Stocks API URL (required)").Required().StringVar(&cfg.APIKey)
	app.Flag("n-days", "The number of days for stock-ticker to dial back when fetching stocks data (required)").Required().IntVar(&cfg.NDays)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
