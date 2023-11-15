package database

import (
	"database/sql"
	"log"
	"os"
	"web_service_gin/schemas"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func InicializeDb() (*gorm.DB, error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASS") + "@tcp(127.0.0.1:3306)/recordings?parseTime=true&loc=Local"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
		return nil, err
	}
	err = gormDB.AutoMigrate(&schemas.Album{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
