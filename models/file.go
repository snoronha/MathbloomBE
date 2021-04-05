package models

import (
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	Path     string `json:"path" sql:"type:varchar(128);unique_index"`
	TicketId uint   `json:"ticketId"`
}
