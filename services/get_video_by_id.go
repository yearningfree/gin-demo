package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// GetVideo
type GetVideo struct {
}

// GetVideoByID 通过ID查找视频
func (getVideo *GetVideo) GetVideoByID(id string) serialzer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serialzer.Response{
			Status: 404,
			Msg:    "该视频不存在",
			Error:  err.Error(),
		}
	}

	video.AddView()

	return serialzer.Response{
		Data: serialzer.BuildVideo(video),
	}
}
