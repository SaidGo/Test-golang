
DB_DSN := "postgres://postgres:1987@localhost:8088/tasksdb?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции: make migrate-new NAME=tasks
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение всех миграций
migrate:
	$(MIGRATE) up

# Откат миграции (на один шаг назад)
migrate-down:
	$(MIGRATE) down

# Запуск сервера
run:
	go run main.go
