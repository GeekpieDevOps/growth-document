package main

import (
	"encoding/json"
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

type DefaultResponse struct {
	Code    int64                  `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type RequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	ID       string `json:"id"`
}


type UpdateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

// 定义基本的 User 模型
type User struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}



func UpdateUser(c *gin.Context,db *gorm.DB) {
	id := c.Param("id")

	//检查用户是否存在
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"code":500,"message": "User not found"})
		return
	}

	//对用户的email, password, nickname进行更新
	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	err := c.ShouldBindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code":500,"message": "Invalid request data"})
		return
	}

	result = db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code":500,"message": "Failed to update user"})
		return
	}

	response := UpdateResponse{
		Code: 200,
		Email:    requestData.Email,
		Password: requestData.Password,
		ID:       user.Data.ID,
	}
	response.Message = "User updated successfully"

	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")

    // 检查用户是否存在
    var user User
    result := db.First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"code":500,"message": "User not found"})
        return
    }

    // 删除用户
    result = db.Delete(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code":500,"message": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"code":200,"message": "User deleted successfully"})
}



func main() {
	// 初始化 Gin
	r := gin.Default()
	// 初始化 Gorm
	// var db *gorm.DB // 由于后面的代码中使用的是简短模式 := ，此处的定义是冗余的
	// 临时数据库启动命令(Docker): docker run -id --name=postgres-test -v postgre-data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=123456 -e LANG=C.UTF-8 postgres
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" //postgres 是 PostgreSQL 的默认用户名和数据库名称
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})                                                                     //"   "应该填链接，由于我不知道postgres的链接形式，所以这里只是一个示例
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")
	// 准备创建表的SQL语句
	createTableStmt := `
    	CREATE TABLE IF NOT EXISTS userss (
        	id SERIAL PRIMARY KEY,
        	email TEXT NOT NULL,
        	password TEXT NOT NULL,
        	nickname TEXT NOT NULL
        )
    `
	// 执行SQL语句
	result := db.Exec(createTableStmt)
	if result.Error != nil {
		fmt.Println("执行SQL语句失败:", result.Error)
		return
	}
	fmt.Println("表创建成功")
	tableName := "userss"

	// 查询列名
	var columns []string
	result = db.Raw("SELECT column_name FROM information_schema.columns WHERE table_name = ?", tableName).Scan(&columns)
	if result.Error != nil {
		fmt.Println("查询列名失败:", result.Error)
		return
	}
	fmt.Println("表", tableName, "的列名:")
	for _, columnName := range columns {
		fmt.Println(columnName)
	}

	// 自动迁移数据库模式
	db.AutoMigrate(&User{})
	// 设置路由
	r.Use(corsMiddleware())
	// r.POST("/your-endpoint", YourHandler)
	// 运行服务

	// Hello World HTTP Endpoint -> JSON
	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/api/v1/login", func(c *gin.Context) {
		Login(c, db)
	})

	r.POST("/api/v1/register", func(c *gin.Context) {
		Register(c, db)
	})

	r.PUT("/api/v1/update", func(c *gin.Context) {//将UpdateUser函数与相应的HTTP请求方法和路径进行关联
		UpdateUser(c, db)
	})

	r.DELETE("/api/v1/delete", func(c *gin.Context) {//将DeleteUser函数与相应的HTTP请求方法和路径进行关联
		DeleteUser(c, db)
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

/*func YourHandler(c *gin.Context) {
	// 读取请求体中的数据
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// 处理读取请求体错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// 打印请求体中的数据
	fmt.Println("string(body)")
	fmt.Println(body)
	// 其他处理逻辑...
}写的貌似有问题，暂时弃用*/

//http://127.0.0.1:18080/api/v1/login
