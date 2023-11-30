package controller

import (
	"assignment-two/config"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = config.ConnectGorm()
}
