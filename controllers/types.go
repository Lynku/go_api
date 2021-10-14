package controllers

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

// GET /types
func TypesList(c *gin.Context) {
	var data []models.PageType
	models.DB.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

//GET /type_fileds/:id
func TypeFileds(c *gin.Context) {
	result := []models.PageTypeOptions{}

	query := models.DB.Model(&models.PageTypeOptions{}).
		Where("page_type_id = ?", c.Param("id")).
		Scan(&result)

	if err := query.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
