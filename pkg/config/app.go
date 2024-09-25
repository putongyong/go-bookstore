package config

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect function will attempt to connect to MySQL, retrying until it succeeds
func Connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Retry logic for MySQL connection
	for {
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second) // Retry every 5 seconds
		} else {
			fmt.Println("Database connected successfully!")
			break
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
