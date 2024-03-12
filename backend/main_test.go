package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: split into small sections for different endpoints
// TODO: use temporary database
func TestEverything(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	router := setupRouter(db)

	t.Run("POST /api/v1/register 500", func(t *testing.T) {
		req, err := http.NewRequest(
			"POST", "/api/v1/register",
			strings.NewReader(``))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}
