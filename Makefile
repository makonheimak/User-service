# Переменные для миграций
DB_DSN_USERS := "postgres://postgres:yourpassword@localhost:5433/usersdb?sslmode=disable"
DB_DSN_TASKS := "postgres://postgres:yourpassword@localhost:5434/tasksdb?sslmode=disable"

MIGRATE_USERS := migrate -path ./migrations -database $(DB_DSN_USERS)
MIGRATE_TASKS := migrate -path ./migrations -database $(DB_DSN_TASKS)

# Миграции для Users
migrate-users-up:
	$(MIGRATE_USERS) up

migrate-users-down:
	$(MIGRATE_USERS) down

# Миграции для Tasks
migrate-tasks-up:
	$(MIGRATE_TASKS) up

migrate-tasks-down:
	$(MIGRATE_TASKS) down

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Запуск всех сервисов
up:
	docker-compose up --build

# Остановка всех сервисов
down:
	docker-compose down

# Просмотр логов
logs:
	docker-compose logs -f

.PHONY: migrate-users-up migrate-users-down migrate-tasks-up migrate-tasks-down migrate-new up down logs