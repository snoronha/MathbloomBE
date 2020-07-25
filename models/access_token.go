package models

import (
	_ "github.com/jinzhu/gorm"
)

type AccessToken struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	UserId        uint   `json:"userId"`
	AccessToken   string `json:"accessToken" gorm:"type:varchar(256);unique_index"`
	TokenId       string `json:"tokenId" gorm:"type:varchar(2048)"`
	ExpiresAt     uint64 `json:"expiresAt"`
	ExpiresIn     uint64 `json:"expiresIn"`
	FirstIssuedAt uint64 `json:"firstIssuedAt"`
}
