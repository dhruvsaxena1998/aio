package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}

		log.Println("Connected to database")

		err = db.AutoMigrate(
			&models.User{},
			&models.Permission{},
			&models.RoleGroup{},
			&models.Habit{},
			&models.Completion{},
		)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
			return
		}

		log.Println("Database migration completed")
	})
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Panic("Database connection not initialized. Call database.Init() first.")
	}
	return db
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Error getting underlying SQL database:", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed")
		}
		db = nil
	}
}
