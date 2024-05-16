package router

import "github.com/gin-gonic/gin"

func Init() {
	// Creates a router without any middleware by default
	r := gin.Default()

	// Define the route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on
}
