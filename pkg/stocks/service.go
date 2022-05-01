package stocks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Service interface {
	GetDataFor(symbol string, daysSince int) (*TickerData, error)
}

type service struct {
	apiKey   string
	apiURL   string
	function string

	cache map[string]cacheData
}

func NewService(apiURL, apiKey string) Service {
	return &service{
		apiURL:   apiURL,
		function: defaultFunction,
		apiKey:   apiKey,

		cache: map[string]cacheData{},
	}
}

func (s *service) WithFunction(function string) {
	s.function = function
}

func (s *service) GetDataFor(symbol string, daysSince int) (*TickerData, error) {
	queryStr := fmt.Sprintf("apikey=%s&function=%s&symbol=%s", s.apiKey, s.function, symbol)
	u, err := url.Parse(s.apiURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryStr
	var jsonData []byte
	if ok := validateCache(s.cache, symbol); ok {
		jsonData = s.cache[symbol].data
	} else {
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		jsonData = body
		s.cache[symbol] = cacheData{
			data:        jsonData,
			lastUpdated: time.Now(),
		}
	}

	days := getDates(time.Now(), daysSince)
	stockResp, err := decodeResponse(jsonData)
	if err != nil {
		return nil, err
	}

	return processStocksData(stockResp, days)
}

func getDates(t time.Time, daysSince int) []string {
	var days []string
	curr := t
	for i := daysSince; i > 0; {
		if isWeekend(curr) {
			curr = curr.Add(-time.Hour * 24)
			continue
		}
		days = append(days, curr.Format(timeFormat))
		curr = curr.Add(-time.Hour * 24)
		i--
	}
	return days
}

func processStocksData(input *TickerData, days []string) (*TickerData, error) {
	output := TickerData{}
	output.Metadata = input.Metadata
	output.OutputData.NDays = len(days)
	if input.Timeseries != nil {
		output.Timeseries = map[string]TimeseriesData{}
		var sum float64
		for day, data := range input.Timeseries {
			if contains(days, day) {
				output.Timeseries[day] = data
				if c, err := strconv.ParseFloat(data.Close, 32); err == nil {
					sum += c
				} else {
					return nil, err
				}
			}

		}
		output.OutputData.CloseAvg = fmt.Sprintf("%f", sum/float64(len(output.Timeseries)))
	}

	return &output, nil
}
