package controllers

import (
	"net/http"
	"web/models"

	"github.com/gin-gonic/gin"
)

// GET /pages/:id
func AdminListPages(c *gin.Context) {
	var data []models.Page
	models.DB.Where("page_type_id = ?", c.Param("id")).Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /page/:id
func AdminGetPage(c *gin.Context) {

	type PageData struct {
		Title      string
		Content    string
		PageTypeId int
	}
	type ResultPage struct {
		PageData
		Fields []map[string]interface{}
	}

	data := ResultPage{}
	page := PageData{}

	q := models.DB.Model(&models.Page{}).Select("pages.page_type_id, pages.title, pages.content").
		Where("pages.id = ? AND pages.status = 1", c.Param("id")).Scan(&page)

	if err := q.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	cf := []map[string]interface{}{}
	models.DB.Model(&models.PageTypeOptions{}).
		Select("page_cfs.value, page_type_options.type, page_type_options.name").
		Joins("left join page_cfs on page_type_options.id = page_cfs.page_type_options_id").
		Where("page_cfs.page_id = ? OR page_type_options.page_type_id = ? ", c.Param("id"), page.PageTypeId).
		Scan(&cf)

	data.Title = page.Title
	data.Content = page.Content
	data.PageTypeId = page.PageTypeId
	data.Fields = cf

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /data/:id
// Get data by ID
func DataById(c *gin.Context) {
	var data models.Page

	if err := models.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// POST /data/new
// Add new data
func NewData(c *gin.Context) {
	var input models.DataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create data
	data := models.Page{
		PageTypeId: input.PageTypeId,
		Author:     input.Author,
		Title:      input.Title,
		Content:    input.Content,
		Status:     input.Status,
	}
	models.DB.Save(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// PUT /data/:id
// Update data
func UpdateData(c *gin.Context) {
	var data models.Page

	if err := models.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.DataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&data).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// DELETE /data/:id
// Delete data
func DeleteData(c *gin.Context) {
	// Get model if exist
	var data models.Page
	if err := models.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
