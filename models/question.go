package models

import (
	"time"
)

type Question struct {
	ID         uint   `form:"id" json:"id" gorm:"primaryKey"`
	UserId     uint   `json:"userId"`
	Question   string `form:"question" json:"question" gorm:"type:varchar(1024)"`
	IsAnswered bool   `form:"isAnswered" json:"isAnswered" gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
