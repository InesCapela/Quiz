package routes

import (
	"Project_2021_PSRS/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllGames(c *gin.Context) {
	controllers.GetAllGames(c)
}

func CreateGame(c *gin.Context) {
	controllers.CreateGame(c)
}

func UpdateGame(c *gin.Context) {
	controllers.UpdateGame(c)
}

func DeleteGame(c *gin.Context) {
	controllers.DeleteGame(c)
}

