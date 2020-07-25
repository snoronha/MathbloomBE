package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
    ID                 uint    `json:"id" gorm:"primary_key"`
    Email              string  `json:"email" sql:"type:varchar(64);unique_index"`
    GoogleId           string  `json:"googleId" gorm:"type:varchar(32)"`
    GivenName          string  `json:"givenName" gorm:"type:varchar(32)"`
    FamilyName         string  `json:"familyName" gorm:"type:varchar(32)"`
    Name               string  `json:"name" gorm:"type:varchar(64)"`
    ImageUrl           string  `json:"imageUrl" gorm:"type:varchar(256)"`
    // Quantity           uint    `json:"quantity"` // from OrderItem
}
