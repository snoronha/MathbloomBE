package controllers

import (
	"MathbloomBE/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /item_details/:item_id
func GetItemDetail(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	itemId, _ := strconv.ParseInt(c.Param("item_id"), 10, 64)
	var itemDetail models.ItemDetail
	db.Where("us_item_id = ?", itemId).First(&itemDetail)
	c.JSON(http.StatusOK, gin.H{"item_detail": itemDetail})
}
