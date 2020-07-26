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
func UpsertAccessTokenWithEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		var accessToken models.AccessToken
		err := c.BindJSON(&accessToken)
		if err != nil {
			log.Print(err)
		}
		// Upsert into access_tokens 
		db.Where(models.AccessToken{UserId: user.ID}).
			Assign(models.AccessToken{
				AccessToken: accessToken.AccessToken,
				TokenId: accessToken.TokenId,
				ExpiresAt: accessToken.ExpiresAt,
				ExpiresIn: accessToken.ExpiresIn,
				FirstIssuedAt: accessToken.FirstIssuedAt,
			}).
			FirstOrCreate(&accessToken)
		c.JSON(http.StatusOK, gin.H{"error": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "no user exists for email: " + email})
	}
}
