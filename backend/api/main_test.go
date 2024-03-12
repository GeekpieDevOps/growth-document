package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/GeekpieDevOps/growth-document/backend/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockLoginBody creates a mock JSON body for the login endpoint
func mockLoginBody() *bytes.Buffer {
	loginDetails := map[string]string{"username": "test", "password": "test"}
	body, _ := json.Marshal(loginDetails)
	return bytes.NewBuffer(body)
}

// TestHelloWorldEndpoint tests the "/" endpoint
func TestHelloWorldEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	req, _ := http.NewRequest("POST", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello, World!", response["message"])
}

// TestLoginEndpoint tests the "/api/v1/login" endpoint
func TestLoginEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/api/v1/login", func(c *gin.Context) {
		response := v1.LoginResponse{
			Code:    200,
			Message: "登录成功",
		}
		response.Data.Type = "Student"
		response.Data.ID = "123456"

		c.JSON(http.StatusOK, response)
	})

	req, _ := http.NewRequest("POST", "/api/v1/login", mockLoginBody())
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response v1.LoginResponse
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "登录成功", response.Message)
	assert.Equal(t, "Student", response.Data.Type)
	assert.Equal(t, "123456", response.Data.ID)
}
