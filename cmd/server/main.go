package main

import (
	"log"
	"os"

	"github.com/makonheimak/user-service/internal/database"
	transportgrpc "github.com/makonheimak/user-service/internal/transport/grpc"
	"github.com/makonheimak/user-service/internal/user/orm"
	"github.com/makonheimak/user-service/internal/user/repository"
	"github.com/makonheimak/user-service/internal/user/service"
)

func main() {
	// Установка переменных окружения по умолчанию для локального запуска
	if os.Getenv("DATABASE_DSN") == "" {
		os.Setenv("DATABASE_DSN", "host=localhost user=postgres password=yourpassword dbname=usersdb port=5433 sslmode=disable")
	}
	// Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}

	// Автомиграция
	if err := db.AutoMigrate(&orm.User{}); err != nil {
		log.Fatalf("❌ Could not migrate database: %v", err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewService(repo)

	// Запуск gRPC
	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("Users gRPC server error: %v", err)
	}
}
