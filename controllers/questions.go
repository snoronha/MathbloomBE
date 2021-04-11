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
	var files []models.File
	db := c.MustGet("db").(*gorm.DB)
	// Check if user with email exists
	email := c.Param("email")
	var user models.User
	db.Where("email = ?", email).First(&user)
	if user.ID > 0 {
		db.Where("user_id = ?", user.ID).Find(&questions)
		// iterate over questions to get question.ID
		// retrieve answers <-> question.IDs, files <-> question.FileTicketId
		questionIds := []uint{}
		fileTicketIds := []uint{}
		for _, question := range questions {
			questionIds = append(questionIds, question.ID)
			if question.FileTicketId > 0 {
				fileTicketIds = append(fileTicketIds, question.FileTicketId)
			}
		}
		// log.Print(questionIds)
		db.Where("question_id IN (?)", questionIds).Find(&answers)
		for _, answer := range answers {
			if answer.FileTicketId > 0 {
				fileTicketIds = append(fileTicketIds, answer.FileTicketId)
			}
		}
		if len(fileTicketIds) > 0 {
			db.Where("ticket_id IN (?)", fileTicketIds).Find(&files)
		}
		c.JSON(http.StatusOK, gin.H{"questions": questions, "answers": answers, "files": files})
	} else {
		c.JSON(http.StatusOK, gin.H{"questions": questions, "answers": answers, "files": files})
	}
}

// POST /question/email/:email
// Derive userId from :email and upsert into questions
func UpsertQuestionWithEmail(c *gin.Context) {
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

		var question models.Question
		if err := c.ShouldBind(&question); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + " with email: " + email})
			return
		}

		// Upsert into questions
		question.UserId = user.ID
		if ticketId > 0 { // need to check if question.FileTicketId is nil?
			question.FileTicketId = ticketId
		}
		// log.Printf("question: %+v\n", question)
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

func GetNextTicket(db *gorm.DB, ticketType string) uint {
	// update ticketId for ticketType and fetch ticketId
	var ticket models.Ticket
	db.Where("ticket_type = ?", ticketType).First(&ticket)
	if ticket.TicketId > 0 {
		db.Model(&ticket).Update("ticket_id", ticket.TicketId+1)
		return ticket.TicketId
	}
	return 0
}
