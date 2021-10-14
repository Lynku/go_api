package main

import (
	"net/http"
	"web/controllers"
	"web/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	// // gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/root")

	// // hit "localhost:8080/root/*
	authorized.GET("/users", controllers.UsersList)
	authorized.GET("/types", controllers.TypesList)
	authorized.GET("/type_fileds/:id", controllers.TypeFileds)

	authorized.GET("/pages", controllers.AdminListPages)
	authorized.GET("/page/:id", controllers.AdminGetPage)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello"})
	})

	r.Use(CORSMiddleware())
	r.Run("127.0.0.5:80")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
