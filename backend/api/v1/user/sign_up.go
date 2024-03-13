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

// SignUpRequest contains fields for an sign up request.
type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignUpResponse specifies response fields for a successful sign up.
type SignUpResponse struct {
	UUID uuid.UUID `json:"uuid" binding:"required,uuid"`
}

// SignIn is the handler for /api/v1/user/sign_up.
func SignUp(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignUpRequest

		// Parse request payload
		if err := c.ShouldBindJSON(&req); err != nil {
			// Inform the client about invalid fields
			if v, ok := err.(validator.ValidationErrors); ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"fields": lo.Map(v, func(f validator.FieldError, _ int) string {
						return f.Field()
					}),
				})
				return
			}

			// Can't handle this, so abort
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var user models.User

		// Check if the user has already signed up
		result := db.Where("email = ?", req.Email).First(&user)
		if result.Error == nil {
			// User already exists
			c.AbortWithStatus(http.StatusConflict)
			return
		} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Can't handle this, so abort
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		// Generate nonce with CSRNG
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
