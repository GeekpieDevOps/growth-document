package user

import (
	"crypto/rand"
	"errors"
	"net/http"

	"github.com/GeekpieDevOps/growth-document/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignUpResponse struct {
	UUID uuid.UUID `json:"uuid" binding:"required,uuid"`
}

func SignUp(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignUpRequest
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
		result := db.Where("email = ?", req.Email).First(&user)
		if result.Error == nil {
			c.AbortWithStatus(http.StatusConflict)
			return
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		nonce := make([]byte, 16)
		_, err := rand.Read(nonce)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user = models.User{
			UUID:     uuid.New(),
			ID:       "",
			Email:    req.Email,
			Password: req.Password,
			Nickname: "",
			Nonce:    nonce,
		}

		result = db.Create(&user)
		if result.Error != nil {
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		c.JSON(http.StatusOK, SignUpResponse{
			UUID: user.UUID,
		})
	}
}
