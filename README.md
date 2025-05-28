# Go Task API (CRUD-сервер на Go)

Простой HTTP-сервер на Go, реализующий полный CRUD (Create, Read, Update, Delete) для задач.

---

## 📦 Установка и запуск

```bash
git clone https://github.com/SaidGo/Test-golang.git
cd Test-golang
go run main.go
```

---

## 🔗 Доступные маршруты

### ➕ POST /task
Создаёт новую задачу.

**Запрос:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"text":"вынести мусор"}' http://localhost:8080/task
```

**Ответ:**
```json
{
  "id": 1,
  "text": "вынести мусор",
  "done": false
}
```

---

### 📖 GET /task
Получает список всех задач.

```bash
curl http://localhost:8080/task
```

---

### ✏ PATCH /task/{id}
Обновляет текст или статус задачи.

**Пример:**
```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"done":true}' http://localhost:8080/task/1
```

**Можно обновить также `text`:**
```bash
curl -X PATCH -H "Content-Type: application/json" -d '{"text":"вынести бутылки"}' http://localhost:8080/task/1
```

---

### ❌ DELETE /task/{id}
Удаляет задачу по ID.

```bash
curl -X DELETE http://localhost:8080/task/1
```

Ответ: `204 No Content`

---

## 📌 Структура задачи (Task)

```json
{
  "id": 1,
  "text": "вынести мусор",
  "done": false
}
```

---

## 👨‍💻 Автор

- **SaidGo**
- GitHub: [@SaidGo](https://github.com/SaidGo)