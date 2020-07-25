package controllers

import (
	"MathbloomBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /access_token/:user_id
func GetAccessToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var accessToken models.AccessToken
	db.Where("user_id = ?", userId).First(&accessToken)
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

// POST /user
func InsertAccessToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var accessToken models.AccessToken
	err := c.BindJSON(&accessToken)
	if err != nil {
		log.Print(err)
	}
	db.Create(&accessToken)
	fail := db.NewRecord(accessToken) // check if insert succeeded
	if !fail {
		c.JSON(http.StatusOK, gin.H{"error": ""})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "accessToken exists"})
	}

}
