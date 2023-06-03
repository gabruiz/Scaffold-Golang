package database

import (
	"encoding/json"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Url      string `json:"database_url"`
	Port     string `json:"database_port"`
	User     string `json:"database_user"`
	Password string `json:"database_password"`
	Name     string `json:"database_name"`
}

var config databaseConfig

func SetParameters(fileData []byte) {
	err := json.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}
}

func Connection() *gorm.DB {
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Url + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	return db
}
