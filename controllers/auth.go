package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
	"web/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /login
// Login
func login(c *gin.Context) {
	var user models.Users
	var token models.Tokens

	password, _ := bcrypt.GenerateFromPassword([]byte(c.Param("password")), 14)
	var checkUser = models.DB.Where("email = ? AND password", c.Param("email"), string(password)).First(&user).Error
	if err := checkUser; err != nil {
		models.DB.Where("id = ?", user).Delete(&token)
		token := newToken(user)
		c.JSON(http.StatusOK, gin.H{"data": "Bearer " + token.Token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User or Password!"})
	}
}

// GET /logout
// Get user by ID
func logout(c *gin.Context) {
	var token models.Tokens
	reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	models.DB.Where("token = ?", reqToken).Delete(&token)

	c.JSON(http.StatusOK, gin.H{"data": "done"})
}

func newToken(u models.Users) models.Tokens {
	var token models.Tokens
	b := make([]byte, 4)
	rand.Read(b)
	var newToken = base64.StdEncoding.EncodeToString(b)
	models.DB.Create(&models.Tokens{
		UserId: u.ID,
		Token:  newToken,
	})
	return token
}
