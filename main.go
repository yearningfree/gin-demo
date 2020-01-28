package main

import (
	"myownsite/server/cache"
	"myownsite/server/model"
	"myownsite/server/route"
)

func main()  {

	// 连接数据库
	model.InitDatabase()
	// 链接redis
	cache.InitRedis()

	r := route.InitRoute()
	r.Run() // 默认监听8080端口
}