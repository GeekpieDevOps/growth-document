package user

import (
	"errors"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/api/v1/constant"
	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type SignOutRequest struct{
	Token string `json:"token" binding:"required,jwt"`
	UUID string `json:"uuid" binding "required,uuid"`
}
func SignOut(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){

		//解析，并检查字段是否合法。此处操作同sign_in sign_up
		var req SignOutRequest
		if err:=c.ShouldBindJSON(&req);err!=nil{
			if v,ok:=err.(validator.ValidationErrors);ok{
				c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
					"fields":lo.Map(v,func(f validator.FieldError,_ int)string{
						return f.Field()
					}),
				})
				return
			}
			c.AbortWithStatus(http.StatusBadRequest)
		}

		//检查是否找到了用户的登录令牌
		var token models.Token
		result:=db.Where("token = ? AND uuid = ?",req.Token,req.UUID).First(&token)
		if result.Error!=nil{
			if errors.Is(result.Error,gorm.ErrRecoredNotFound){
				//未找到用户的登录令牌
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}else{
				//其他的未知错误
				c.AbortWithError(http.StatusInternalServerError,result.Error)
				return
			}
		}

		//删除相应的令牌，实现登出操作
		result:=db.Where("token = ? AND uuid = ?",req.Token,req.UUID).Delete(&token)
		if result.Error!=nil{
			//删除操作出错
			c.AbortWithError(http.StatusInternalServerError,result.Error)
			return
		}
		if result.RowsAffected==0{
			//操作后数据库行数未收到影响，即未找到匹配的记录
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK,gin.H{
			//返回登出成功
			"message":"Sign out succesfully!"
		})
	}
}