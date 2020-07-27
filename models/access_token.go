package models

import (
	"github.com/jinzhu/gorm"
)

type AccessToken struct {
	gorm.Model
	UserId        uint   `json:"userId"`
	AccessToken   string `json:"accessToken" gorm:"type:varchar(256);unique_index"`
	TokenId       string `json:"tokenId" gorm:"type:varchar(3072)"`
	ExpiresAt     uint64 `json:"expiresAt"`
	ExpiresIn     uint64 `json:"expiresIn"`
	FirstIssuedAt uint64 `json:"firstIssuedAt"`
}
