package stocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateCache(t *testing.T) {
	t.Run("Test Non-existent entry in cache map", func(t *testing.T) {
		cache := map[string]cacheData{}
		cache["sampleKey"] = cacheData{}
		assert.False(t, validateCache(cache, "sampleKey2"))
	})

	t.Run("Fresh entry in cache map", func(t *testing.T) {
		cache := map[string]cacheData{}
		t1, err := time.Parse(timeFormat, "2021-01-01")
		Now = func() time.Time {
			return t1
		}
		assert.Nil(t, err)
		cache["sampleKey"] = cacheData{
			data:        []byte{},
			lastUpdated: t1,
		}
		assert.True(t, validateCache(cache, "sampleKey"))
	})
}
