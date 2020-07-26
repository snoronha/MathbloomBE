package models

import (
	_ "github.com/jinzhu/gorm"
)

type Problem struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	Guid          string  `json:"guid" gorm:"type:varchar(40);unique_index"`
	UserId        uint    `json:"userId"`
    Specs         string  `json:"specs" sql:"type:varchar(512)"`
    Answer        string  `json:"answer" gorm:"type:varchar(64)"`
	Attempt       string  `json:"attempt" gorm:"type:varchar(64)"`
}
