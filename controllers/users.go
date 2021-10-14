package controllers

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func UsersList(c *gin.Context) {
	var users []models.Users
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /user/:id
// Get user by ID
func User(c *gin.Context) {
	var user models.Users

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
