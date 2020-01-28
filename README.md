# 上手做一个自己的网站练习
## 加油，奥利给！

Gin+gorm+redis+mysql

#### api功能

	// 检测后段服务器是否运行畅通
	GET    /api/v1/ping
	
	// 用户注册
	POST   /api/v1/user/register
	
	// 用户登陆
	POST   /api/v1/user/login
	
	// 查看用户信息
	GET    /user/info
	
	// 用户登出
	DELETE /user/logout
	
	// 创建视频
	POST   /api/v1/videoCreate
	
	// 获取某个视频
	GET    /api/v1/video/:id
	
	// 获取全部视频
	GET    /api/v1/videos
	
	// 修改视频
	PUT    /api/v1/video/:id
	
	// 删除视频
	DELETE /api/v1/video/:id