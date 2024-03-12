package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SignUp(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			if v, ok := err.(validator.ValidationErrors); ok {
				c.JSON(http.StatusBadRequest,
					lo.Map(v, func(f validator.FieldError, _ int) string {
						return f.Field()
					}),
				)
				return
			}

			c.String(http.StatusBadRequest, "")
			return
		}

		c.String(http.StatusNotImplemented, "")
	}
}
