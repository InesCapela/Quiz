package routes

import (
	"Project_2021_PSRS/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllQuestions(c *gin.Context) {
	controllers.GetAllQuestions(c)
}

func CreateQuestion(c *gin.Context) {
	controllers.CreateQuestion(c)
}

func UpdateQuestion(c *gin.Context) {
	controllers.UpdateQuestion(c)
}

func DeleteQuestion(c *gin.Context) {
	controllers.DeleteQuestion(c)
}
