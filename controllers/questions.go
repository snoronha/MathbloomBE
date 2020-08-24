package controllers

import (
	"MathbloomBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /question/id/:id
func GetQuestionById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var question models.Question
	db.Where("id = ?", id).First(&question)
	c.JSON(http.StatusOK, gin.H{"question": question})
}

// GET /questions/email/:email
// Get all questions asked by user/email
func GetQuestionsByEmail(c *gin.Context) {
	var questions []models.Question
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		db.Where("user_id = ?", user.ID).Find(&questions)
		c.JSON(http.StatusOK, gin.H{"questions": questions})
	} else {
		c.JSON(http.StatusOK, gin.H{"questions": questions})
	}
}

// POST /question/email/:email
// Derive userId from :email and upsert into questions
func UpsertQuestionWithEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		var question models.Question
		err := c.BindJSON(&question)
		if err != nil {
			log.Print(err)
		}
		// Upsert into questions
		question.UserId = user.ID
		if question.ID > 0 {
			db.Assign(models.Question{
				UserId:     user.ID,
				Question:   question.Question,
				IsAnswered: question.IsAnswered,
			}).
				FirstOrCreate(&question)
		} else {
			db.Create(&question)
		}
		c.JSON(http.StatusOK, gin.H{"error": "", "id": question.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "no user exists for email: " + email})
	}
}
