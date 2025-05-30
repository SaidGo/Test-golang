package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Модель задачи
type Task struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=1987 dbname=tasksdb port=8088 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}

// POST /tasks
func createTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	task := Task{Text: input.Text, IsDone: input.IsDone}
	db.Create(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GET /tasks
func listTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// PATCH /tasks/{id}
func updateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	var input struct {
		Text   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if input.Text != nil {
		task.Text = *input.Text
	}
	if input.IsDone != nil {
		task.IsDone = *input.IsDone
	}

	db.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DELETE /tasks/{id}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&Task{}, id).Error; err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTask(w, r)
		case http.MethodGet:
			listTasks(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			updateTask(w, r)
		case http.MethodDelete:
			deleteTask(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
