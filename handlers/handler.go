package handlers

import (
	database "web_service_gin/db"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeHandler() {
	var err error
	db, err = database.InicializeDb()
	if err != nil {
		panic(err)
	}
}
