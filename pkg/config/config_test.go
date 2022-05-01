package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
	cfg := NewConfig()
	args := []string{
		"--symbol=GOOG",
		"--api-key=sw34kjnsd23123asf2",
		"--api-url=http://betavantage.co/query",
		"--n-days=3",
	}
	assert.Nil(t, cfg.ParseFlags(args))
	assert.Equal(t, 3, cfg.NDays)
	assert.Equal(t, "GOOG", cfg.Symbol)
	assert.Equal(t, "sw34kjnsd23123asf2", cfg.APIKey)
	assert.Equal(t, "http://betavantage.co/query", cfg.APIURL)
}
