package controllers

import (
	"MathbloomBE/models"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // To generate random file names
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
	var answers []models.Answer
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		db.Where("user_id = ?", user.ID).Find(&questions)
		// iterate over questions to get question.ID
		// retrieve answers with those question.IDs
		questionIds := []uint{}
		for _, question := range questions {
			questionIds = append(questionIds, question.ID)
		}
		// log.Print(questionIds)
		db.Where("question_id IN (?)", questionIds).Find(&answers)
		c.JSON(http.StatusOK, gin.H{"questions": questions, "answers": answers})
	} else {
		c.JSON(http.StatusOK, gin.H{"questions": questions, "answers": answers})
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
		if err := c.ShouldBind(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " with email: " + email})
			return
		}
		// Upsert into questions
		question.UserId = user.ID
		log.Printf("question: %+v\n", question)
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

// POST /question/attachment/email/:email
// Derive userId from :email and upsert into questions
func UpsertQuestionAttachmentWithEmail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]
		for _, file := range files {
			// Retrieve file information
			extension := filepath.Ext(file.Filename)
			// Generate random file name for the new uploaded file so it doesn't override the old file with same name
			newFileName := uuid.New().String() + extension
			// log.Printf("Saving %s as %s\n", filepath.Base(file.Filename), newFileName)

			// The file is received, so let's save it
			if err := c.SaveUploadedFile(file, "/Users/macbook/go/src/MathbloomBE/uploads/"+newFileName); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
				return
			}
		}

		// File saved successfully. Return proper result
		c.JSON(http.StatusOK, gin.H{"message": "Your file has been successfully uploaded"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "no user exists for email: " + email})
	}
}
