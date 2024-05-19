package router

import "github.com/gin-gonic/gin"

func Init() {
	// Creates a router without any middleware by default
	r := gin.Default()

	// Initialize routes
	InitializeRoutes(r)

	r.Run() // listen and serve on
}
