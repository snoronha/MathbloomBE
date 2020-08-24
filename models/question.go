package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	UserId     uint   `json:"userId"`
	Question   string `json:"question" gorm:"type:varchar(1024)"`
	IsAnswered bool   `json:"isAnswered" gorm:"default:false"`
}
