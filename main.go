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

	router.GET("/user/:user_id", controllers.GetUser)
	router.POST("/user", controllers.InsertUser)
	router.GET("/access_token/:user_id", controllers.GetAccessToken)
	router.POST("/access_token", controllers.InsertAccessToken)
	// router.GET("/item/:id", controllers.FindItem) // new
	// router.GET("/items", controllers.FindItems) // new
	// router.GET("/items/search", controllers.SearchItems)
	// router.GET("/item_detail/:item_id", controllers.GetItemDetail)

	router.Run(":8081")
}
