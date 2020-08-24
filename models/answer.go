package models

import (
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	QuestionId uint   `json:"questionId"`
	UserId     uint   `json:"userId"`
	Answer     string `json:"answer" gorm:"type:varchar(1024)"`
	IsAccepted bool   `json:"isAccepted" gorm:"default:false"`
}
