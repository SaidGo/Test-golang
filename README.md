# 📌 Go Task API с PostgreSQL + Users

REST API-сервер на Go с использованием Echo, GORM, PostgreSQL и OpenAPI генерацией.  
Реализованы CRUD-операции для `tasks` и `users`. Используется `strict-server` подход.

---

## 🚀 Как запустить

1. Убедитесь, что PostgreSQL установлен и запущен.
2. Создайте базу данных:
```sql
CREATE DATABASE tasksdb;
```
3. Клонируйте репозиторий и запустите сервер:
```bash
git clone https://github.com/SaidGo/Test-golang.git
cd Test-golang
go run ./cmd/app
```

Сервер доступен по адресу: `http://localhost:8080`

---

## ⚙️ Команды Makefile

```bash
make run                      # запуск сервера
make lint                     # проверка кода линтером
make lint-fix                 # автоисправление ошибок
make gen                      # генерация API по OpenAPI
make migrate-new NAME=users  # новая миграция
make migrate-up               # применить миграции
make migrate-down             # откат последней миграции
```

---

## 🧪 Линтер

Установка:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Запуск:
```bash
golangci-lint run --out-format=colored-line-number
```

---

## 📂 Структура проекта

```
.
├── cmd/app            # main.go
├── internal
│   ├── database       # инициализация БД
│   ├── tasksService   # бизнес-логика и модель Tasks
│   ├── userService    # бизнес-логика и модель Users
│   ├── handlers       # HTTP-обработчики (tasks, users)
│   └── web
│       ├── tasks      # сгенерированный код API tasks
│       └── users      # сгенерированный код API users
├── openapi            # спецификация OpenAPI и конфиги
└── migrations         # SQL-миграции
```

---

## 🔗 Примеры API-запросов (curl)

### ➕ POST /tasks
```bash
curl -X POST -H "Content-Type: application/json" -d '{"task":"Пример", "is_done":false}' http://localhost:8080/tasks
```

### 📖 GET /tasks
```bash
curl http://localhost:8080/tasks
```

### ✏ PATCH /tasks/{id}
```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"task":"Обновлено","is_done":true}' http://localhost:8080/tasks/1
```

### ❌ DELETE /tasks/{id}
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 👥 Работа с Users (через Postman)

- **GET /users** — получить список пользователей  
- **POST /users** — создать пользователя  
- **PATCH /users/{id}** — обновить пользователя  
- **DELETE /users/{id}** — удалить пользователя

📌 Используется генерация через oapi-codegen и strict-handlers.

---

## 🛠 Подключение к PostgreSQL

- Host: `localhost`  
- Port: `8088`  
- User: `postgres`  
- Password: `1987`  
- DB name: `tasksdb`

---

## 📦 Зависимости

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [OpenAPI](https://swagger.io/specification/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [golangci-lint](https://golangci-lint.run/)

---

## 👨‍💻 Автор

**SaidGo**  
GitHub: [@SaidGo](https://github.com/SaidGo)