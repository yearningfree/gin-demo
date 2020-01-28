package serialzer

import "gin-demo/model"

// Video 视频序列化器
type Video struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Avatar      string `json:"avatar"`
	View        uint64 `json:"view"` // 点击数
	CreatedAt   int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {

	return Video{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		URL:         item.URL,
		Avatar:      item.Avatar,
		//View:      item.View(),
		CreatedAt: 0,
	}
}

// BuildVideos 序列化视频
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
