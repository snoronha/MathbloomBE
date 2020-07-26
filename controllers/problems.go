package controllers

import (
	"MathbloomBE/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /problem/id/:id
func GetProblemById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var problem models.Problem
	db.Where("id = ?", id).First(&problem)
	c.JSON(http.StatusOK, gin.H{"problem": problem})
}

// GET /problems/email/:email
// Get all problems attempted/solved by user/email
func GetProblemsByEmail(c *gin.Context) {
	var problems []models.Problem
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		db.Where("user_id = ?", user.ID).Find(&problems)
		c.JSON(http.StatusOK, gin.H{"problems": problems})
	} else {
		c.JSON(http.StatusOK, gin.H{"problems": problems})
	}
}

// POST /problem/email/:email
// Derive userId from :email and upsert into problems
func UpsertProblemWithEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		var problem models.Problem
		err := c.BindJSON(&problem)
		if err != nil {
			log.Print(err)
		}
		// Upsert into access_tokens 
		db.Where(models.Problem{Guid: problem.Guid}).
			Assign(models.Problem{
				UserId: user.ID,
				Guid: problem.Guid,
				Specs: problem.Specs,
				Answer: problem.Answer,
				Attempt: problem.Attempt,
			}).
			FirstOrCreate(&problem)
		c.JSON(http.StatusOK, gin.H{"error": ""})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "no user exists for email: " + email})
	}
}
