package model

import (
	"github.com/jinzhu/gorm"
	"gin-demo/cache"
	"strconv"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title       string // 视频标题
	Description string // 视频简介
	URL         string // 视频链接
	Avatar      string // 视频封面
}

// View 点击数
func View(video *Video) uint64 {
	countStr, err := cache.RedisCli.Get(cache.VideoViewKey(video.ID)).Result()
	if err != nil {
		panic(err)
		return 0
	}
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加点击数
func (video *Video) AddView() {
	click := cache.VideoViewKey(video.ID)

	// 增加视频点击数
	cache.RedisCli.Incr(click)

	// 增加排行榜点击数
	cache.RedisCli.ZIncrBy(cache.DailyRank, 1, strconv.Itoa(int(video.ID)))

}
