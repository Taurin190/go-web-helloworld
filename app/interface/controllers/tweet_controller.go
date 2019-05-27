package controller

import (
	"strconv"

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

func (controller *TweetController) Detail(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	print(id)
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
