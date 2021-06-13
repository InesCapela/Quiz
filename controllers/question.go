package controllers

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllQuestions /**
func GetAllQuestions(c *gin.Context) {
	var questions []model.Question

	services.OpenDatabase()
	services.Db.Select([]string{"question","id"}).Find(&questions)
	defer services.Db.Close()

	for i, question := range questions {
		var options []model.Options
		services.Db.Where("question_id = ?", question.ID).Find(&options)
		//services.Db.Select([]string{"question_id","option"}).Find(&options)
		questions[i].Options = options
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": questions})
}

// CreateQuestion /**
func CreateQuestion(c *gin.Context) {
	var question model.Question

	if err := c.BindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check syntax!"})
		return
	}
	services.Db.Save(&question)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Create successful!", "resourceId": question})
}

// UpdateQuestion /**
func UpdateQuestion(c *gin.Context) {
	var question model.Question

	id := c.Param("id")
	services.Db.First(&question, id)

	if question.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Place not found!"})
		return
	}

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	services.Db.Save(question)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Update succeeded!"})
}

func DeleteQuestion(c *gin.Context) {
	var question model.Question

	id := c.Param("id")
	services.Db.First(&question, id)

	if question.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&question)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}