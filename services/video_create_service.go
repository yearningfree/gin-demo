package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// VideoCreateService 视频投稿数据表单
type VideoCreateFrom struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Description string `form:"description" json:"info" binding:"max=3000"`
	URL         string `form:"url" json:"url"`
	Avatar      string `form:"avatar" json:"avatar"`
}

// VideoCreate 视频投稿
func (vc *VideoCreateFrom) VideoCreate() serialzer.Response {
	video := model.Video{
		Title:       vc.Title,
		Description: vc.Description,
		URL:         vc.URL,
		Avatar:      vc.Avatar,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serialzer.Response{
			Status: 5003,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}

	return serialzer.Response{
		Data: serialzer.BuildVideo(video),
	}
}
