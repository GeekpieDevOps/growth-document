package user

import (
    
)

type SignOutRequest struct {
	UUID string `json:"uuid" binding:"required,uuid"`
	token string `json:"token" binding:"required,jwt"`
}

func SignOut(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignOutRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			if v, ok := err.(validator.ValidationErrors); ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"fields": lo.Map(v, func(f validator.FieldError, _ int) string {
						return f.Field()
					}),
				})
				return
			}

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// 获取请求中的用户信息
		userID := c.GetString("userID")

		// 在数据库中查找对应的用户
		var user models.User

		result := db.First(&user, "uuid = ?", userID)
		if result.Error != nil {
			// 用户不存在
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				// 处理错误，中止请求
				c.AbortWithError(http.StatusInternalServerError, result.Error)
				return
			}
		}

        // 检查令牌是否存在并有效
        tokenString := req.token
        if tokenString == "" {//调用路径暂不清楚
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(""), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

		// 更新用户的登录状态为未登录
		result = db.Model(&user).Update("logged_in", false)
		if result.Error != nil {
			// 处理错误，中止请求
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		// 删除用户的令牌
		result = db.Delete(&models.Token{}, "uuid = ?", userID)
		if result.Error != nil {
			// 处理错误，中止请求
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User signed out successfully",
		})
	}
}