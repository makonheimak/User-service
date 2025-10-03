package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Используем переменную окружения или fallback на имя контейнера
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		// Для Docker-сети используем имя контейнера
		dsn = "host=postgres-users user=postgres password=yourpassword dbname=usersdb port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	log.Println("✅ Connected to Users database")
	return DB, nil
}
