package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRank 每日排行key
	DailyRank = "rank:daily"
)

// VideoViewKey 视频点击数的key
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
