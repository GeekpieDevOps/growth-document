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

type GetRequest struct{
	UUID string `json:"uuid" binding:"required,uuid"`
	ID string `json:"id" `
	Password string `json "password`
}

func GetUser(db *gorm.DB)func(c *gin.Context){
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

		//对用户的令牌进行解析
		parseToken,err:=jwt.ParseWithClaims(req.Token,&CustomClaims{},func(parseToken *jwt.Token)(i interface{},err error){
			return Nonce,nil
		})
		if err!=nil{
			return 
		}

		//检查是否找到了用户的登录令牌
		var token models.Token
		resultToken:=db.Where("uuid = ?",req.UUID).First(&token)
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
/**/
		//检查是否找到了用户
		var user models.User
		result:=db.Where("uuid = ? AND id = ? AND password = ",req.UUID,req.ID,req.Password).First(&user)
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

		//返回用户信息
		c.JSON(http.StatusOK,gin.H{
			"code":"200",
			"message":"Sccessfully found",
			"data":gin.H{
				"邮箱":user.Email,
			},
		})
	}
}