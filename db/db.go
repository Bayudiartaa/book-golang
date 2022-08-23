package db

import (
	"fmt"

	"github.com/Bayudiartaa/book-golang/configs"
	"github.com/Bayudiartaa/book-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ConnectDatabase() (db *gorm.DB, err error) {
	dbHost, dbUser, dbPwd, dbName, dbSsl := configs.GetDBInfo()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=Asia/Jakarta", dbHost, dbUser, dbPwd, dbName, dbSsl)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// database, err := gorm.Open("sqlite3", "test.db")

	// if err != nil {
	// 	log.Panic("Failed to connect to database!")
	// 	return
	// }

	db.AutoMigrate(&models.Book{})

	return db, err
}
