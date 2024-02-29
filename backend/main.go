package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

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
	r.Use(corsMiddleware())
	r.POST("/your-endpoint", YourHandler)

	// 运行服务

	// Hello World HTTP Endpoint -> JSON
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/api/v1/login", func(c *gin.Context) {
		response := LoginResponse{
			Code:    200,
			Message: "登录成功",
		}
		response.Data.Type = "Student"
		response.Data.ID = "123456"
		var body = make([]byte, 0)
		c.Request.Body.Read(body)
		fmt.Print(body)

		c.JSON(http.StatusOK, response)
	})
	// 运行服务
	err = r.Run(":18080")
	if err != nil {
		log.Fatal(err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func YourHandler(c *gin.Context) {
	// 读取请求体中的数据
	//body,
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// 处理读取请求体错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// 打印请求体中的数据
	fmt.Println("string(body)")
	fmt.Println(string(body))
	// 其他处理逻辑...
}

//http://127.0.0.1:18080/api/v1/login环境
