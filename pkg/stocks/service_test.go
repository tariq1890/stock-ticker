package stocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetDays(t *testing.T) {
	t1, err := time.Parse(timeFormat, "2022-03-03")
	assert.Nil(t, err)
	dates := getDates(t1, 5)
	assert.Equal(t, 5, len(dates))
	assert.ElementsMatch(t, []string{"2022-03-03", "2022-03-02", "2022-03-01", "2022-02-28", "2022-02-25"}, dates)
}

func TestProcessStockData(t *testing.T) {
	inp, err := decodeResponse([]byte(testData))
	assert.Nil(t, err)

	out, err := processStocksData(inp, []string{"2022-03-03", "2022-03-02", "2022-03-01", "2022-02-28", "2022-02-25"})
	assert.Nil(t, err)
	assert.NotNil(t, out.OutputData)
	assert.NotNil(t, out.Metadata)
	assert.NotNil(t, out.Timeseries)
	assert.Equal(t, out.OutputData.CloseAvg, "15.470000")
	assert.Equal(t, out.OutputData.NDays, 5)
	assert.Equal(t, out.Metadata.Symbol, "FORG")
	assert.Equal(t, 5, len(out.Timeseries))
}
