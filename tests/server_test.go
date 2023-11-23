// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Set the Gin to Test mode
	gin.SetMode(gin.TestMode)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a request to pass to our handler
	req, _ := http.NewRequest("GET", "/hello", nil)

	// Create a Gin router with the handler we want to test
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, test!"})
	})

	// Perform the request and check the status code
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expected := `{"message":"Hello, test!"}`
	assert.Equal(t, expected, w.Body.String())
}