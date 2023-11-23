/**
 * Author: Pavel
 * File: server.go
 */

package main

import (
	"github.com/cutebex/sportvenue-server/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// For Server Connection Test
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	// listen and serve on 0.0.0.0:8080
	// on windows "localhost:8080"
	// can be overriden with the PORT env var

	venueController := controllers.VenueController{}

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/venue/:keyword", venueController.Get)
	}

	router.Run()
}
