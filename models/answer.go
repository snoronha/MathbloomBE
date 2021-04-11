package models

import (
	"time"
)

type Answer struct {
	ID           uint   `form:"id" json:"id" gorm:"primaryKey"`
	QuestionId   uint   `form:"questionId" json:"questionId"`
	UserId       uint   `json:"userId"`
	Answer       string `form:"answer" json:"answer" gorm:"type:varchar(1024)"`
	FileTicketId uint   `form:"fileTicketId" json:"fileTicketId"`
	IsAccepted   bool   `form:"isAccepted" json:"isAccepted" gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
