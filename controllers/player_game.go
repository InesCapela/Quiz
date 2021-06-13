package controllers

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddPlayerGame /**
func AddPlayerGame(c *gin.Context) {

	services.OpenDatabase()

	username, _ := c.Get("username")
	user := model.Users{}
	user.Username = username.(string)

	var game model.Game
	id := c.Param("id")

	services.Db.Select("id").First(&game, id)

	game.Players = append(game.Players, user)

	services.Db.Save(&game)

	services.Db.Close()

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": game})
}

// CheckAnswer /**
func CheckAnswer(c *gin.Context) {

	type response struct {
		QuestionID uint   `json:"question_id" binding:"required"`
		Response   string `json:"question_answer" binding:"required"`
	}

	services.OpenDatabase()
	defer services.Db.Close()

	// id do game
	id := c.Param("id")
	var game model.Game
	services.Db.Where("id = ?", id).Find(&game)

	if game.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invalid game id!"})
	}

	// My user
	username, _ := c.Get("username")
	username = username.(string)

	// Players
	var players []model.Users
	services.Db.Model(&game).Association("Players").Find(&players)

	found := false
	for _, player := range players {
		if player.Username == username {
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Não estas a participar neste jogo!"})
	}

	// Resposta
	var resposta response
	err := c.BindJSON(&resposta)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
	}

	// Questions
	var questions []model.Question
	services.Db.Model(&game).Association("Question").Find(&questions)

	for _, question := range questions {
		if question.ID == resposta.QuestionID {
			if question.Answer == resposta.Response {
				game.Score++
				services.Db.Save(&game)
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Assertou!", "score": game.Score})
			} else {
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Tente novamente!", "score": game.Score})
			}
		}
	}

}
