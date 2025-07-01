# 📌 Go Task API с PostgreSQL + Users

REST API-сервер на Go с использованием **Echo**, **GORM**, **PostgreSQL** и генерации кода по **OpenAPI**.  
Реализованы CRUD-операции для `tasks` и `users`. Используется строгий подход к маршрутам (`strict-server`), код OpenAPI сгенерирован через `oapi-codegen`.

🔥 Поля вида `*_at` (`created_at`, `updated_at`) исключены из JSON-ответов на уровне сериализации.

---

## 🚀 Быстрый старт

1. Установите PostgreSQL и создайте базу:
   ```sql
   CREATE DATABASE tasksdb;
   ```

2. Клонируйте репозиторий и запустите сервер:
   ```bash
   git clone https://github.com/SaidGo/Test-golang.git
   cd Test-golang
   go run ./cmd/app
   ```

Сервер доступен по адресу: [`http://localhost:8080`](http://localhost:8080)

---

## ⚙️ Makefile команды

```bash
make run                      # запуск сервера
make lint                     # запуск линтера
make lint-fix                 # автоисправление ошибок
make gen                      # генерация API-кода из openapi/*.yaml
make migrate-new NAME=users  # создание новой миграции
make migrate-up               # применение всех миграций
make migrate-down             # откат последней миграции
```

---

## 🧪 Линтинг

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
├── cmd/app              # main.go
├── internal/
│   ├── database         # инициализация БД
│   ├── tasksService     # бизнес-логика Tasks
│   ├── userService      # бизнес-логика Users
│   ├── handlers         # HTTP-обработчики (tasks, users)
│   └── web/
│       ├── tasks        # сгенерированный код API tasks
│       └── users        # сгенерированный код API users
├── openapi/             # OpenAPI-спеки и конфиги генерации
├── migrations/          # SQL-миграции
└── Makefile             # команды запуска и сборки
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

📌 Используется генерация через `oapi-codegen` и строгие хендлеры.

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
