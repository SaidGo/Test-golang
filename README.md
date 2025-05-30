
# 📌 Go Task API с PostgreSQL

Простой REST API-сервер на Go с использованием GORM и PostgreSQL.  
Реализованы все CRUD-операции (создание, чтение, обновление, удаление задач) с использованием базы данных.

---

## 🚀 Как запустить

1. Убедитесь, что PostgreSQL установлен и запущен.
2. Создайте базу данных:
```sql
CREATE DATABASE tasksdb;
```
3. Клонируйте репозиторий:
```bash
git clone https://github.com/SaidGo/Test-golang.git
cd Test-golang
go run main.go
```

Сервер будет доступен по адресу:  
`http://localhost:8080`

---

## 🔗 Примеры API-запросов

### ➕ POST /tasks  
Создание новой задачи:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"task":"Проверить почту", "is_done":false}' http://localhost:8080/tasks
```

### 📖 GET /tasks  
Получение списка всех задач:
```bash
curl http://localhost:8080/tasks
```

### ✏ PATCH /tasks/{id}  
Обновление задачи по ID:
```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"is_done":true}' http://localhost:8080/tasks/1
```

### ❌ DELETE /tasks/{id}  
Удаление задачи по ID (мягкое удаление):
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## 🛠 Подключение к PostgreSQL

Сервер подключается к базе данных PostgreSQL:

- Host: `localhost`  
- Port: `8088`  
- User: `postgres`  
- Password: `1987`  
- DB name: `tasksdb`

---

## 👨‍💻 Автор

**SaidGo**  
GitHub: [@SaidGo](https://github.com/SaidGo)
