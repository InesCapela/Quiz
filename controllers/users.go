package controllers

import (
	"Project_2021_PSRS/model"
	"Project_2021_PSRS/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllUsers /**
func GetAllUsers(c *gin.Context) {
	var users []model.Users
	username := c.Param("username")

	services.Db.Find(&users, username)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

// GetUserByID /**
func GetUserByID(c *gin.Context) {
	var user model.Users
	id := c.Param("id")

	services.Db.First(&user, id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

// UpdateUser /**
func UpdateUser(c *gin.Context) {
	var user model.Users

	id := c.Param("id")
	services.Db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found!"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Check request!"})
		return
	}

	// Update settings
	c.Set("username", user.Username)
	c.Set("isAdmin", user.IsAdmin)

	services.Db.Save(user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Update succeeded!"})
}

// DeleteUser /**
func DeleteUser(c *gin.Context) {
	var user model.Users

	id := c.Param("id")
	services.Db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}

	services.Db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete succeeded!"})
}
