# Go Task API

Простой HTTP-сервер на Go с поддержкой GET и POST запросов.  
Позволяет сохранять задачу и возвращать приветствие с её содержимым.

## 📦 Установка и запуск

```bash
git clone https://github.com/SaidGo/Test-golang.git
cd Test-golang
go run main.go
```

## 🔗 Примеры запросов

### POST /
Сохраняет задачу:

```bash
curl -X POST -H "Content-Type: application/json" -d "{"task":"вынести мусор"}" http://localhost:8080
```

Ответ:
```
Задача сохранена: вынести мусор
```

### GET /
Получить приветствие:

```bash
curl http://localhost:8080
```

Ответ:
```
hello, вынести мусор
```

## 👨‍💻 Автор

- **SaidGo**
- GitHub: [@SaidGo](https://github.com/SaidGo)
