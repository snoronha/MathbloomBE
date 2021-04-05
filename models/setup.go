// models/setup.go

package models

import (
	"MathbloomBE/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db := util.GetDB()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&AccessToken{})
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Answer{})
	db.AutoMigrate(&Ticket{})
	db.AutoMigrate(&File{})

	return db
}
