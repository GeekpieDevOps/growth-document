package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义基本的 User 模型
type User struct {
	gorm.Model
	Name  string
	StuID string
}

func main() {
	// 初始化 Gin
	r := gin.Default()

	// 初始化 Gorm
	// var db *gorm.DB // 由于后面的代码中使用的是简短模式 := ，此处的定义是冗余的

	// 临时数据库启动命令(Docker): docker run -id --name=postgres-test -v postgre-data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=123456 -e LANG=C.UTF-8 postgres
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" //postgres 是 PostgreSQL 的默认用户名和数据库名称

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //"   "应该填链接，由于我不知道postgres的链接形式，所以这里只是一个示例
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	// 自动迁移数据库模式
	db.AutoMigrate(&User{})

	// 设置路由

	// Hello World HTTP Endpoint -> JSON
	r.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Database Endpoint -> JSON
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// 运行服务
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
