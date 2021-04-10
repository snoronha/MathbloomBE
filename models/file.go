package models

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Path     string `json:"path" sql:"type:varchar(128);unique_index"`
	Url      string `json:"url" sql:"type:varchar(128)"`
	TicketId uint   `json:"ticketId"`
}
