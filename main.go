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
	//services.Db.DropTableIfExists(&model.Options{})
	//services.Db.DropTableIfExists(&model.Game{})

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
		auth.POST("/register", routes.RegisterUser)
		auth.PUT("/refresh_token", services.AdminAuthorizationRequired(), routes.RefreshToken)
	}
}

func questionsRoutes(router *gin.Engine) {

	back := router.Group("/user")
	{
		back.PUT("/:id", services.UserAuthorizationRequired(), routes.UpdateUser)
		back.GET("/", services.UserAuthorizationRequired(), routes.GetAllUsers)
		back.GET("/:id", services.AdminAuthorizationRequired(), routes.GetUserByID)
		back.DELETE("/:id", services.AdminAuthorizationRequired(), routes.DeleteUser)
	}

	back = router.Group("/game")
	{
		back.POST("/", services.AdminAuthorizationRequired(), routes.CreateGame)
		back.PUT("/:id",  services.AdminAuthorizationRequired(), routes.UpdateGame)
		back.DELETE("/:id",  services.AdminAuthorizationRequired(), routes.DeleteGame)
		back.GET("/", services.AdminAuthorizationRequired(), routes.GetAllGames)
		back.GET("/:id", services.UserAuthorizationRequired(), routes.GetUserFromGame)
		back.POST("/join/:id", services.UserAuthorizationRequired(), routes.AddPlayerGame)
		back.POST("/play/:id", services.UserAuthorizationRequired(), routes.CheckAnswer)
	}

	back = router.Group("/question")
	{
		back.POST("/", services.AdminAuthorizationRequired(), routes.CreateQuestion)
		back.PUT("/:id",  services.AdminAuthorizationRequired(), routes.UpdateQuestion)
		back.DELETE("/:id",  services.AdminAuthorizationRequired(), routes.DeleteQuestion)
		back.GET("/", services.AdminAuthorizationRequired(), routes.GetAllQuestions)
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
