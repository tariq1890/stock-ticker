package stocks

import (
	"time"
)

type cacheData struct {
	data        []byte
	lastUpdated time.Time
}

func validateCache(cache map[string]cacheData, key string) bool {
	now := time.Now()
	if c, ok := cache[key]; ok {
		if now.Sub(c.lastUpdated) < (time.Hour * 24) {
			return true
		}
	}
	return false
}
