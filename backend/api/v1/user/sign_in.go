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

// SignInRequest contains fields for an sign in request.
type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignInResponse specifies response fields for a successful sign in.
type SignInResponse struct {
	Token string    `json:"token" binding:"required,jwt"`
	UUID  uuid.UUID `json:"uuid" binding:"required,uuid"`
}

// SignIn is the handler for /api/v1/user/sign_in.
func SignIn(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignInRequest

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

		// Check if provided information are correct
		result := db.Where("email = ? AND password = ?", req.Email, req.Password).First(&user)
		if result.Error != nil {
			// User doesn't exist
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				// Can't handle this, so abort
				c.AbortWithError(http.StatusInternalServerError, result.Error)
				return
			}
		}

		// Allocate token ID
		id := uuid.New()

		// Create a session token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Subject:  constant.JWTSubjectAuthentication.String(),
			Audience: jwt.ClaimStrings{user.UUID.String()},
			ID:       id.String(),
		})

		// Sign it with user's nonce
		signedToken, err := token.SignedString(user.Nonce)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.SetCookie("token", signedToken, 3600, "/", "", false, true)

		result = db.Create(&models.Token{
			ID:    id,
			UUID:  user.UUID,
			Nonce: user.Nonce,
			Token: signedToken,//由密钥生成的签名
		})
		if result.Error != nil {
			// Can't handle this, so abort
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		c.JSON(http.StatusOK, SignInResponse{
			Token: signedToken,
			UUID:  user.UUID,
		})
	}
}
