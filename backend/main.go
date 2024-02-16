package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 初始化 Gin
	r := gin.Default()

	// 初始化 Gorm
	var db *gorm.DB
	db, err := gorm.Open(postgres.Open(" ? "), &gorm.Config{}) //"   "应该填链接，由于我不知道postgres的链接形式，所以这里只是一个示例
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	// 设置路由

	// 运行服务

}
