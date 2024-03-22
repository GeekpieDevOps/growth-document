func UpdateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req UpdateUserRequest

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

		// Get the user ID from the request context or session
		userID := getUserIDFromContextOrSession(c) // You need to implement this function

		// Retrieve the user from the database
		var user models.User
		result := db.First(&user, userID)
		if result.Error != nil {
			// User doesn't exist
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			} else {
				// Can't handle this, so abort
				c.AbortWithError(http.StatusInternalServerError, result.Error)
				return
			}
		}

		// Update the user's email and password
		user.Email = req.Email
		user.Password = req.Password

		// Save the updated user to the database
		result = db.Save(&user)
		if result.Error != nil {
			// Can't handle this, so abort
			c.AbortWithError(http.StatusInternalServerError, result.Error)
			return
		}

		c.Status(http.StatusOK)
	}
}