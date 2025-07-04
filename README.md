# 📌 Go Task API с SQLite + Users

REST API-сервер на Go с использованием **Echo**, **GORM**, **SQLite** и генерации кода по **OpenAPI**.  
Реализованы CRUD-операции для `tasks` и `users`, включая привязку задач к конкретным пользователям.  
Используется строгий подход к маршрутам (`strict-server`), код сгенерирован через `oapi-codegen`.

🔥 Поля `*_at` (`created_at`, `updated_at`) исключены из JSON-ответов.

---

## 🚀 Быстрый старт

1. Убедитесь, что у вас установлен Go ≥ 1.23 и SQLite.

2. Клонируйте репозиторий и установите зависимости:
   ```bash
   git clone https://github.com/SaidGo/Test-golang.git
   cd Test-golang
   go mod tidy
   ```

3. Запустите сервер:
   ```bash
   make run
   ```

Сервер будет доступен на [`http://localhost:8080`](http://localhost:8080)

---

## ⚙️ Команды Makefile

```bash
make run                        # запуск сервера
make gen                        # генерация API-кода из openapi/*.yaml (users.yaml использует tasks.yaml через import-mapping)
make migrate-new NAME=xxx      # создание новой миграции
make migrate-up                 # применение всех миграций
make migrate-down               # откат последней миграции
make lint                       # проверка линтером
make lint-fix                   # автоисправление
make tidy                       # очистка/обновление зависимостей
```

---

## 🧪 Линтинг

Установка:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Запуск:
```bash
make lint
```

---

## 🗃️ Миграции

Установка:
```bash
# для Windows (пример):
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.windows-amd64.zip -o migrate.zip
# распаковать и добавить migrate.exe в PATH
```

---

## 📂 Структура проекта

```
.
├── cmd/app                 # main.go — точка входа
├── internal/
│   ├── database            # инициализация SQLite
│   ├── handlers            # HTTP-хендлеры (tasks, users)
│   ├── tasksService        # бизнес-логика и модели задач
│   ├── userService         # бизнес-логика и модели пользователей
│   └── web/
│       ├── tasks           # сгенерированный OpenAPI код (tasks)
│       └── users           # сгенерированный OpenAPI код (users)
├── openapi/                # OpenAPI-спецификации и oapi-codegen.yaml
├── migrations/             # SQL-миграции для SQLite
├── tasks.db                # локальная БД
├── Makefile                # команды запуска и сборки
└── go.mod / go.sum         # зависимости
```

---

## 🔗 Примеры запросов (curl)

### ➕ POST /tasks
```bash
curl -X POST http://localhost:8080/tasks   -H "Content-Type: application/json"   -d '{"task":"Пример", "is_done":false, "user_id":1}'
```

### 📖 GET /tasks
```bash
curl http://localhost:8080/tasks
```

### ✏ PATCH /tasks/{id}
```bash
curl -X PATCH http://localhost:8080/tasks/1   -H "Content-Type: application/json"   -d '{"task":"Обновлено", "is_done":true}'
```

### ❌ DELETE /tasks/{id}
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 👥 Users (через Postman или curl)

- `GET /users`
- `POST /users`
- `PATCH /users/{id}`
- `DELETE /users/{id}`
- `GET /users/{id}/tasks` — задачи пользователя

---

## 📦 Зависимости

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/)
- [OpenAPI 3](https://swagger.io/specification/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [golangci-lint](https://golangci-lint.run/)
- [migrate](https://github.com/golang-migrate/migrate)

---

## 👨‍💻 Автор

**SaidGo**  
GitHub: [@SaidGo](https://github.com/SaidGo)