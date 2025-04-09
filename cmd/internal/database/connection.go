package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD")+"",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	log.Println("Connected to database")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Habit{},
		&models.Completion{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
