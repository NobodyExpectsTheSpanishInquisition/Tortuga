package Infrastructure

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	r := gin.Default()

	registerRoutes(r)

	return r
}

func registerRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
