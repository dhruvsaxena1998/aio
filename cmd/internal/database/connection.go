package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dhruvsaxena1998/aio/cmd/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func handleMigrations(database *gorm.DB) error {
	if err := database.AutoMigrate(
		&models.User{},
		&models.Permission{},
		&models.RoleGroup{},
		&models.Habit{},
		&models.Completion{},
	); err != nil {
		log.Panicf("Failed to migrate database: %v", err)
		return err
	}
	log.Println("Database migration completed")
	return nil
}

func Init() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			getEnv("DB_HOST"), getEnv("DB_USER"), getEnv("DB_PASSWORD"),
			getEnv("DB_NAME"), getEnv("DB_PORT"), getEnv("DB_SSLMODE"),
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})
		if err != nil {
			log.Panicf("Failed to connect to database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Panicf("Failed to get SQL DB from GORM: %v", err)
		}

		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
		sqlDB.SetConnMaxIdleTime(10 * time.Minute)

		log.Println("Connected to database and configured connection pool")

		if err := handleMigrations(db); err != nil {
			log.Panicf("Failed to handle migrations: %v", err)
		}
	})
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Panic("Database not initialized. Call database.Init() first.")
	}
	return db
}

func CloseDB() {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting SQL DB: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	} else {
		log.Println("Database connection closed")
	}
	db = nil
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Panicf("Environment variable %s is not set", key)
	}
	return val
}
