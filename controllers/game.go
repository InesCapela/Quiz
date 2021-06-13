package controllers

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

// GetAllGames /**
func GetAllGames(c *gin.Context) {
	var game []model.Game

	services.Db.Select("id").Find(&game)

	if len(game) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": game})
}

// CreateGame /**
func CreateGame(c *gin.Context) {
	nPerguntas := 3
	var game model.Game
	var allQuestion []model.Question
	var myQuestions []model.Question

	services.Db.Find(&allQuestion)

	rand.Seed(0)
	n := rand.Intn(len(allQuestion) - nPerguntas + 1)
	myQuestions = allQuestion[n : n+nPerguntas]

	for i, question := range myQuestions {
		var options []model.Options
		services.Db.Where("question_id = ?", question.ID).Find(&options)
		myQuestions[i].Options = options
		fmt.Println(options)
	}

	game.Question = myQuestions


	services.Db.Save(&game)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": game})
}

// UpdateGame /**
func UpdateGame(c *gin.Context) {
	var game model.Game

	id := c.Param("id")
	services.Db.First(&game, id)

	if game.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Place not found!"})
		return
	}

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	services.Db.Save(game)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Update succeeded!"})
}

func DeleteGame(c *gin.Context) {
	var game model.Game

	id := c.Param("id")
	services.Db.First(&game, id)

	if game.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&game)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}
