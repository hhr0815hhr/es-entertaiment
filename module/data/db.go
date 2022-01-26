package data

import (
	"es-entertainment/core/database/mysql"

	"github.com/jinzhu/gorm"
)

func WriteDb() *gorm.DB {
	return mysql.GetDB("master")
}

func ReadDb() *gorm.DB {
	return mysql.GetDB("slaver")
}
