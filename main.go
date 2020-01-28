package main

import (
	"gin-demo/cache"
	"gin-demo/model"
	"gin-demo/route"
)

func main()  {
	// gin模式
	//gin.SetMode(gin.ReleaseMode)

	// 连接数据库
	model.InitDatabase()
	// 链接redis
	cache.InitRedis()

	r := route.InitRoute()
	r.Run() // 默认监听8080端口
}