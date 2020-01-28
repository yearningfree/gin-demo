package route

import (
	"github.com/gin-gonic/gin"
	"myownsite/server/api"
	"myownsite/server/middleware"
)

// 初始化路由
func InitRoute() *gin.Engine {
	r := gin.Default()

	// session中间件
	r.Use(middleware.Session("secret"))
	// 跨域中间件
	r.Use(middleware.Cors())

	// api分组
	v1 := r.Group("api/v1/")
	{
		// 检测连接
		v1.GET("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 使用中间件做登录认证
		authorized := r.Group("/")
		// 用户验证
		authorized.Use(middleware.AuthorizedLogin())
		{
			// 需要认证的路由
			// 获取用户信息
			authorized.GET("user/info", api.UserInfo)
			// 用户登出
			authorized.DELETE("user/logout", api.UserLogout)
		}

		// 视频操作
		v1.POST("videoCreate", api.VideoCreate)
		v1.GET("video/:id", api.GetVideoByID)
		v1.GET("videos", api.VideoShow)
		v1.PUT("video/:id", api.VideoUpdate)
		v1.DELETE("video/:id", api.VideoDelete)

	}
	return r
}
