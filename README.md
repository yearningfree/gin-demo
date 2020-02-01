Gin+gorm+redis+mysql


### 视频教程：

[让我们写个G站吧](https://space.bilibili.com/10/channel/detail?cid=78794)

### GO MOD

本项目使用GO MOD管理依赖，适用于go1.11及以后版本。
使用：

	# 模式开启
	export GO111MODULE=on
	# 设置代理
	export GOPROXY=https://mirrors.aliyun.com/proxy
	# 初始化mod文件
	go mod init gin-demo
	# 添加依赖
	go build main.go

[依赖管理](https://github.com/yearningfree/accumulation/blob/master/Mynote/%E6%96%B0%E7%89%B9%E6%80%A7GoModule.md)

### 项目详解

#### 1.网站功能

用户：注册、登录、登出、查看信息；

视频：添加、修改、删除、获取单个或多个视频；

### 2.目录结构

`api`文件夹，用于整合模块具体接口文件；

`cache`文件夹，用于存放redis相关缓存文件；

`middleware`文件夹，用于存放中间件；

`model`文件夹，用于存放数据库模型；

`route`文件夹，用于存放路由信息；

`serialzer`文件夹，用于存通用试图模型和json对象；

`services`文件夹，用于存放具体功能服务。

#### 3.api使用

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

### 鸣谢

[橙卡-bydmm](https://space.bilibili.com/10)

### 加油，奥利给！