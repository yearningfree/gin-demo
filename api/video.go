package api

import (
	"gin-demo/services"
	"github.com/gin-gonic/gin"
)

// VideoCreate 视频投稿
func VideoCreate(c *gin.Context) {
	vCreate := services.VideoCreateFrom{}
	if err := c.ShouldBind(&vCreate); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := vCreate.VideoCreate()
		c.JSON(200, res)
	}
}

// GetVideoByID 根据ID查询视频
func GetVideoByID(c *gin.Context) {
	video := services.GetVideo{}
	res := video.GetVideoByID(c.Param("id"))
	c.JSON(200, res)
}

// VideoShow 视频展示
func VideoShow(c *gin.Context) {
	video := services.VideoShow{}
	if err := c.ShouldBind(&video); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := video.Show()
		c.JSON(200, res)
	}
}

// UpdateVideo 视频更新
func VideoUpdate(c *gin.Context) {
	video := services.VideoUpdateService{}
	if err := c.ShouldBind(&video); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := video.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

// UpdateVideo 视频删除
func VideoDelete(c *gin.Context) {
	video := services.VideoDeleteService{}
	res := video.Delete(c.Param("id"))
	c.JSON(200, res)
}
