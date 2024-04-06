package user

import (
	"errors"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
//	"github.com/go-playground/validator/v10"
//	"github.com/samber/lo"
	"gorm.io/gorm"
//	"github.com/dgrijalva/jwt-go"
)



func SignOut(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){

		//从cookie中获取用户的令牌
		tokenstring, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			return
		}

/*		//对用户的令牌进行解析
		type RegisteredClaims struct{
			Subject string
			Audience string
			ID string
		}

		parseToken,err:=jwt.ParseWithClaims(tokenstring, &RegisteredClaims{}, func(token *jwt.Token)(i interface{},err error){
			return _,nil
		})
		if err!=nil{
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !parseToken.Valid{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}*/

		//检查是否找到了用户的登录令牌
		var token models.Token
		uuid:=c.Param("uuid")

		result:=db.Where("Token = ? AND UUID = ?",tokenstring,uuid).First(&token)
		if result.Error!=nil{
			if errors.Is(result.Error,gorm.ErrRecordNotFound){
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
		resultDelete:=db.Where("Token = ?",tokenstring).Delete(&token)
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

		//检查是否找到了用户
		var user models.User
		tokenResult:=db.Where("UUID = ?",uuid).First(&user)
		if tokenResult.Error!=nil{
			//用户不存在
			if errors.Is(tokenResult.Error,gorm.ErrRecordNotFound){
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}else{
				c.AbortWithError(http.StatusInternalServerError,tokenResult.Error)
				return
			}
		}

		c.JSON(http.StatusOK,gin.H{
			//返回登出成功
			"message":"Sign out succesfully!",
		})
	}
}