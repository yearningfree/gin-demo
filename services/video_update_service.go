package services

import (
	"gin-demo/model"
	"gin-demo/serialzer"
)

// VideoUpdateService 视频更新服务
type VideoUpdateService struct {
	Description string `form:"description" binding:"max=3000"`
}

// Update 更新视频
func (vs *VideoUpdateService) Update(id string) serialzer.Response{
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil{
		return serialzer.Response{
			Status:404,
			Msg:"视频不存在",
			Error:err.Error(),
		}
	}

	video.Description = vs.Description
	err = model.DB.Save(&video).Error
	if err != nil{
		return serialzer.Response{
			Status:5002,
			Msg:"视频保存失败",
			Error:err.Error(),
		}
	}

	return serialzer.Response{
		Data:serialzer.BuildVideo(video),
	}
}
