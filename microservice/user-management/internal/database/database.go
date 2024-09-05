package database

import (
	"fmt"
	"log"
	"user-management-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCONN *gorm.DB

func LoadDB() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.ENV.DB_URL,
		config.ENV.DB_USERNAME,
		config.ENV.DB_PASSWORD,
		config.ENV.DB_NAME,
		config.ENV.DB_PORT)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DBCONN = db

	log.Println("Database running...")
}
