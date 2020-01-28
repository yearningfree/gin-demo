package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 创建一个单例的DB
var DB *gorm.DB

// InitDatabase 初始化数据库
func InitDatabase() {
	db, err := gorm.Open("mysql",
		"root:1234@/ownsite?charset=utf8&parseTime=True&loc=Local")
	//db.logMode(true)

	if err != nil {
		panic(err)
	}

	//if gin.Mode() == "release"{
	//	db.logMode(false)
	//}

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(30)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	autoMigrate()
}

// AutoMigrate 数据库迁移
func autoMigrate() {
	DB.Set("gorm:table_options", "charset=utf8mb4")
	if DB.HasTable(&User{}) == false {
		DB.AutoMigrate(&User{})
	}

	if DB.HasTable(&Video{}) == false {
		DB.AutoMigrate(&Video{})
	}

}
