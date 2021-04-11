package controllers

import (
	"MathbloomBE/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	UPLOADS_DIR := os.Getenv("GOPATH") + "/src/MathbloomBE/uploads"
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

		var ticketId uint = 0
		files := form.File["files"]
		if len(files) > 0 {
			ticketId = GetNextTicket(db, "file") // ticket for ticketType == "file"
		}
		for _, file := range files {
			// Retrieve file information
			extension := filepath.Ext(file.Filename)
			// Generate random file name for the new uploaded file
			// so it doesn't override the old file with same name
			newUuid := uuid.New().String()
			newFileName := newUuid + extension
			hex1 := strings.ToUpper(newUuid[0:2])
			hex2 := strings.ToUpper(newUuid[2:4])
			downloadUrl := "/uploads/" + hex1 + "/" + hex2 + "/" + newFileName
			fullFilePath := UPLOADS_DIR + "/" + hex1 + "/" + hex2 + "/" + newFileName
			// log.Printf("Saving %s as %s\n", filepath.Base(file.Filename), fullFilePath)

			// The file is received, so let's save it
			err := c.SaveUploadedFile(file, fullFilePath)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
				})
				return
			} else {
				// Insert meta data as row in files table
				file := models.File{Path: fullFilePath, Url: downloadUrl, TicketId: ticketId}
				result := db.Create(&file)
				if result.Error != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"message": "File save failed: " + fullFilePath,
					})
					return
				}
			}
		}

		var answer models.Answer
		if err := c.ShouldBind(&answer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " with email: " + email})
			return
		}

		// Upsert into answers
		answer.UserId = user.ID
		if ticketId > 0 { // need to check if question.FileTicketId is nil?
			answer.FileTicketId = ticketId
		}
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
