package routes

import (
	"Project_2021_PSRS/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	controllers.RegisterHandler(c)
}

func GenerateToken(c *gin.Context) {
	controllers.LoginHandler(c)
}

func RefreshToken(c *gin.Context) {
	controllers.RefreshHandler(c)
}
