package controllers

import (
	"MathbloomBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /user/:user_id
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var user models.User
	db.Where("id = ?", userId).First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GET /user/:email
func GetUserByEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// POST /user
func InsertUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err)
	}
	db.Create(&user)
	fail := db.NewRecord(user) // check if insert succeeded
	if !fail {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user exists"})
	}

}
