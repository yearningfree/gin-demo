package services

import (
	"myownsite/server/model"
	"myownsite/server/serialzer"
)

// VideoDeleteService 视频删除服务
type VideoDeleteService struct{
}

// Delete 视频删除
func (vs *VideoDeleteService) Delete(id string) serialzer.Response {
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err !=nil{
		return serialzer.Response{
			Status:404,
			Msg:"视频不存在",
			Error:err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err !=nil{
		return serialzer.Response{
			Status:5003,
			Msg:"视频删除失败",
			Error:err.Error(),
		}
	}
	return serialzer.Response{}
}