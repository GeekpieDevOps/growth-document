package user

import (
	"errors"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
)

type UpdateRequest struct{
	//Token string `json:"token" binding:"required,jwt"`
	Name string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func Update(db *gorm.DB) func(c *gin.Context){
	return func(c *gin.Context){
		var req GetRequest
		//判断请求字段的合法性
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
			return
		}

		//从cookie中获取用户的令牌
		tokenstring, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			return
		}
		
		//对用户的令牌进行解析
		type RegisteredClaims struct{
			Subject: string
			Audience :string
			ID:string
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
		}
		
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

		//检查是否找到了用户
		var user models.User
		result:=db.Where("UUID = ?",uuid).First(&user)
		if result.Error!=nil{
			//用户不存在
			if errors.Is(result.Error,gorm.ErrRecordNotFound){
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}else{
				c.AbortWithError(http.StatusInternalServerError,result.Error)
				return
			}
		}

		switch req.Name{
			case "ID"{
				user.ID=req.Value
			}
			case "Email"{
				user.Email=req.Value
			}
			case "Password"{
				user.Password=req.Value
			}
			case "Nickname"{
				user.Nickname=req.Value
			}
			default:
				c.AbortWithStatus(http.StatusBadRequest)
		}

		c.JSON(http.StatusOK, UpdateResponse{
			"message":"修改成功",
		})

	}
}