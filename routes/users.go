package routes

import (
	"Project_2021_PSRS/controllers"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	controllers.GetAllUsers(c)
}

func GetUserByID(c *gin.Context) {
	controllers.GetUserByID(c)
}
func DeleteUser(c *gin.Context) {
	controllers.DeleteUser(c)
}
func UpdateUser(c *gin.Context) {
	controllers.UpdateUser(c)
}

func GetUserFromGame(c *gin.Context) {
	controllers.GetUserFromGame(c)
}




