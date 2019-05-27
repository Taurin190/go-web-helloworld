package infrastructure

import (
	controllers "../../app/interface/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	tweetController := controllers.GetTweetController(GetDBHandler())
	router.GET("/tweets", func(c *gin.Context) { tweetController.Index(c) })
	router.GET("/tweets/:id", func(c *gin.Context) { tweetController.Detail(c) })

	Router = router
}
