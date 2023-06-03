package dao

import (
	"example.com/main-service/database"
	"gorm.io/gorm"
)

func getDatabase() *gorm.DB {
	return database.Connection()
}

func checkErrors(result *gorm.DB) {
	if result.Error != nil {
		panic(result.Error.Error())
	}
}
