# -----------------------------
# Настройки миграций (SQLite)
# -----------------------------
DB_DSN ?= sqlite3://tasks.db
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создать новую миграцию:
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применить все миграции вверх
migrate-up:
	$(MIGRATE) up

# Откатить последнюю миграцию
migrate-down:
	$(MIGRATE) down

# -----------------------------
# Генерация кода по OpenAPI
# -----------------------------
gen:
	oapi-codegen -config openapi/users.config.yaml openapi/users.yaml
	oapi-codegen -config openapi/tasks.config.yaml openapi/tasks.yaml

# -----------------------------
# Форматирование и линтинг
# -----------------------------
# Форматировать весь код с gofmt
format:
	gofmt -s -w .

# Запустить линтер
lint:
	golangci-lint run --out-format colored-line-number

# Автоматически исправить ошибки линтера
lint-fix:
	golangci-lint run --fix

# Проверка и автоформатирование всего кода
check:
	make tidy format lint

# -----------------------------
# Сборка и запуск приложения
# -----------------------------
run:
	go run cmd/app/main.go

build:
	go build -o app.exe ./cmd/app

# -----------------------------
# Обновление зависимостей
# -----------------------------
tidy:
	go mod tidy
