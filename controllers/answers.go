package controllers

import (
	"MathbloomBE/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /answer/id/:id
func GetAnswerById(c *gin.Context) {
	// TODO: splice in the question in response
	db := c.MustGet("db").(*gorm.DB)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var answer models.Answer
	db.Where("id = ?", id).First(&answer)
	c.JSON(http.StatusOK, gin.H{"answer": answer})
}

// GET /answers//:question_id
// Get all answers for :question_id
func GetAnswersByQuestionId(c *gin.Context) {
	// TODO: splice in the question in response as well
	var answers []models.Answer
	db := c.MustGet("db").(*gorm.DB)
	question_id, _ := strconv.ParseInt(c.Param("question_id"), 10, 64)
	db.Where("question_id = ?", question_id).Find(&answers)
	c.JSON(http.StatusOK, gin.H{"answers": answers})
}

// POST /api/answer/email/:email
// Derive userId from :email and upsert into answers
func UpsertAnswerWithEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		var answer models.Answer
		err := c.BindJSON(&answer)
		if err != nil {
			log.Print(err)
		}
		// Upsert into questions
		answer.UserId = user.ID
		if answer.ID > 0 {
			db.Assign(models.Answer{
				UserId:     user.ID,
				Answer:     answer.Answer,
				QuestionId: answer.QuestionId,
			}).
				FirstOrCreate(&answer)
		} else {
			db.Create(&answer)
		}
		c.JSON(http.StatusOK, gin.H{"error": "", "id": answer.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "no user exists for email: " + email})
	}
}
