package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld(t *testing.T) {
	// 设置 Gin 的模式为测试模式
	gin.SetMode(gin.TestMode)

	// 创建一个 Gin 引擎实例
	r := gin.New()

	// 添加测试的路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// 创建一个 POST 请求
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 将请求发送到路由
	r.ServeHTTP(w, req)

	// 检查状态码
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200; got %v", w.Code)
	}

	// 检查响应体
	expected := `{"message":"Hello, World!"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %v; got %v", expected, w.Body.String())
	}
}
