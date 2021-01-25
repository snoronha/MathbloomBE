package main

import (
	"MathbloomBE/controllers"
	"MathbloomBE/models"
	"MathbloomBE/util"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Check for uploads structure
	util.CreateUploadFileStructure()

	// Setup DB models if not exist
	db := models.SetupModels()

	// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/api/user/id/:user_id", controllers.GetUser)
	router.GET("/api/user/email/:email", controllers.GetUserByEmail)
	router.POST("/api/user", controllers.InsertUser)
	router.GET("/api/access_token/:user_id", controllers.GetAccessToken)
	router.POST("/api/access_token/email/:email", controllers.UpsertAccessTokenWithEmail)
	router.GET("/api/problem/:id", controllers.GetProblemById)
	router.GET("/api/problems/email/:email", controllers.GetProblemsByEmail)
	router.POST("/api/problem/email/:email", controllers.UpsertProblemWithEmail)
	router.GET("/api/question/:id", controllers.GetQuestionById)
	router.GET("/api/questions/email/:email", controllers.GetQuestionsByEmail)
	router.POST("/api/question/email/:email", controllers.UpsertQuestionWithEmail)
	router.POST("/api/question/attachment/email/:email", controllers.UpsertQuestionAttachmentWithEmail)
	router.GET("/api/answer/:id", controllers.GetAnswerById)
	router.GET("/api/answers/:question_id", controllers.GetAnswersByQuestionId)
	router.POST("/api/answer/email/:email", controllers.UpsertAnswerWithEmail)

	router.Run(":8081")
}
