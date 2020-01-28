package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// VideoShow 视频展示
type VideoShow struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// Show 视频展示
func (vs *VideoShow) Show() serialzer.Response {
	videos := []model.Video{}
	total := 0

	if vs.Limit == 0 {
		vs.Limit = 4
	}

	// 获取视频总数
	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serialzer.Response{
			Status: 5002,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	// 分页查询
	if err := model.DB.Limit(vs.Limit).Offset(vs.Start).Find(&videos).Error; err != nil {
		return serialzer.Response{
			Status: 5002,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serialzer.BuildListResponse(serialzer.BuildVideos(videos),uint(total))
}
