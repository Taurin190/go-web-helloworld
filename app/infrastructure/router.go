package infrastructure

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.GET("/tweets", func(c *gin.Context) {})

	Router = router
}
