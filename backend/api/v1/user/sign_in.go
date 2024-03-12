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

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SignIn(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignInRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			if v, ok := err.(validator.ValidationErrors); ok {
				c.AbortWithStatusJSON(http.StatusBadRequest,
					lo.Map(v, func(f validator.FieldError, _ int) string {
						return f.Field()
					}),
				)
				return
			}

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var user models.User
		result := db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				c.AbortWithError(http.StatusInternalServerError, result.Error)
				return
			}
		}

		// TODO
		c.AbortWithStatus(http.StatusNotImplemented)
	}
}
