package main

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/routes"
	"Project_2021_PSRS/services"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var identityKey = "id"

func init() {
	services.OpenDatabase()
	//services.Db.DropTableIfExists(&model.Users{})
	//services.Db.DropTableIfExists(&model.Question{})

	services.Db.AutoMigrate(&model.Users{})
	services.Db.AutoMigrate(&model.Question{})
	services.Db.AutoMigrate(&model.Options{})
	services.Db.AutoMigrate(&model.Game{})

	defer services.Db.Close()
}

func authRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", routes.GenerateToken)
		auth.POST("/register", services.AdminAuthorizationRequired(), routes.RegisterUser)
		auth.PUT("/refresh_token", services.AdminAuthorizationRequired(), routes.RefreshToken)
	}
}

func questionsRoutes(router *gin.Engine) {

	back := router.Group("/admin")
	back.Use(services.AdminAuthorizationRequired())
	{
		back.GET("/users/:id", routes.GetUserByID)
		back.DELETE("/users/:id", routes.DeleteUser)
		back.POST("/question", routes.CreateQuestion)
		back.PUT("/question/:id", routes.UpdateQuestion)
		back.DELETE("/question/:id", routes.DeleteQuestion)
	}

	back = router.Group("/user")
	back.Use(services.UserAuthorizationRequired())
	{
		back.PUT("/users/:id", routes.UpdateUser)
		back.GET("/question", routes.GetAllQuestions)
		back.GET("/users", routes.GetAllUsers)
	}

}
func main() {

	services.FormatSwagger()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authRoutes(router)
	questionsRoutes(router)

	router.Run(":8080")
}
