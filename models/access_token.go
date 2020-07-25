package models

import (
	_ "github.com/jinzhu/gorm"
)

type AccessToken struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	UserId        uint   `json:"userId" gorm:"unique_index"`
	AccessToken   string `json:"accessToken" gorm:"type:varchar(256)"`
	TokenId       string `json:"tokenId" gorm:"type:varchar(1024)"`
	ExpiresAt     uint64 `json:"expiresAt"`
	ExpiresIn     uint64 `json:"expiresIn"`
	FirstIssuedAt uint64 `json:"firstIssuedAt"`
}
