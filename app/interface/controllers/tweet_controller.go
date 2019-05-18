package controller

import (
	"../../../app/interface/database"
	"../../../app/usecase"
	"github.com/gin-gonic/gin"
)

type TweetController struct {
	Interactor usecase.TweetInteractor
}

func GetTweetController(dbHandler database.DBHandler) *TweetController {
	return &TweetController{
		Interactor: usecase.TweetInteractor{},
	}
}

func (controller *TweetController) Index(c Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (controller *TweetController) Show(c Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
