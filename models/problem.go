package models

import (
	_ "github.com/jinzhu/gorm"
)

type Problem struct {
	ID                 string  `json:"id" gorm:"type:varchar(40);unique_index"`
	UserId             uint    `json:"userId"`
    Specs              string  `json:"specs" sql:"type:varchar(256)"`
    Answer             string  `json:"answer" gorm:"type:varchar(64)"`
	Attempt            string  `json:"attempt" gorm:"type:varchar(64)"`
}
