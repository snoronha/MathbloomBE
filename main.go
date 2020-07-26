package main

import (
	"MathbloomBE/controllers"
	"MathbloomBE/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
	// router.GET("/item/:id", controllers.FindItem) // new
	// router.GET("/items", controllers.FindItems) // new
	// router.GET("/items/search", controllers.SearchItems)
	// router.GET("/item_detail/:item_id", controllers.GetItemDetail)

	router.Run(":8081")
}
