package stocks

import (
	"encoding/json"
	"time"
)

func decodeResponse(b []byte) (*TickerData, error) {
	var data TickerData
	err := json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func isWeekend(t time.Time) bool {
	t = t.UTC()
	switch t.Weekday() {
	case time.Friday:
		h, _, _ := t.Clock()
		if h >= 12+10 {
			return true
		}
	case time.Saturday:
		return true
	case time.Sunday:
		h, m, _ := t.Clock()
		if h < 12+10 {
			return true
		}
		if h == 12+10 && m <= 5 {
			return true
		}
	}
	return false
}
