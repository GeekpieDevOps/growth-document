package user

import (
	"errors"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type LogOffRequest struct{
	Token string `json:"token" binding:"required,jwt"`
	UUID string `json:"uuid" binding "required,uuid"`
}
func LogOff(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){

		//解析，并检查字段是否合法。此处操作同sign_in sign_up
		var req LogOffRequest
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
    
		//对用户的令牌进行解析
		parseToken,err:=jwt.ParseWithClaims(req.Token,&CustomClaims{},func(parseToken *jwt.Token)(i interface{},err error){
			return Nonce,nil
		})
		if err!=nil{
			return 
		}

		//检查是否找到了用户的登录令牌
		var token models.Token
		resultToken:=db.Where("token = ? AND uuid = ?",req.Token,req.UUID).First(&token)
		if resultToken.Error!=nil{
			if errors.Is(resultToken.Error,gorm.ErrRecordNotFound){
				//未找到用户的登录令牌
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}else{
				//其他的未知错误
				c.AbortWithError(http.StatusInternalServerError,resultToken.Error)
				return
			}
		}

		//根据提供的uuid查找用户
		var user models.User
		resultUser:=db.Where("uuid = ?",req.UUID).First(&user)
		if resultUser.Error!=nil{
			if errors.Is(resultUser.Error,gorm.ErrRecordNotFound){
				//未找到用户
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}else{
				//其他的未知错误
				c.AbortWithError(http.StatusInternalServerError,resultUser.Error)
				return
			}
		}

		//删除用户
		resultDeleteUser:=db.Where("uuid = ?",req.UUID).Delete(&user)
		if resultDeleteUser.Error!=nil{
			//删除操作出错
			c.AbortWithError(http.StatusInternalServerError,resultDeleteUser.Error)
			return
		}
		if resultDeleteUser.RowsAffected==0{
			//操作后数据库行数未收到影响，即未找到匹配的记录
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		//删除相应的令牌
		resultDelete:=db.Where("token = ? AND uuid = ?",req.Token,req.UUID).Delete(&token)
		if resultDelete.Error!=nil{
			//删除操作出错
			c.AbortWithError(http.StatusInternalServerError,resultDelete.Error)
			return
		}
		if resultDelete.RowsAffected==0{
			//操作后数据库行数未收到影响，即未找到匹配的记录
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		//返回注销成功
		c.JSON(http.StatusOK,gin.H{
			"message":"log off succesfully!",
		})
	}
}