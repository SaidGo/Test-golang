package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/driver/sqlite"
	_ "modernc.org/sqlite"
)

type Task struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Text      string         `json:"task"`
	Done      bool           `json:"is_done"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Dialector{
		DSN:        "tasks.db",
		DriverName: "sqlite", // ключевой момент — указать modernc/sqlite
	}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}
	db.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	var input Task
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if input.Text != "" {
		task.Text = input.Text
	}
	task.Done = input.Done

	db.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	db.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTaskHandler(w, r)
		case http.MethodGet:
			listTasksHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch, http.MethodPut:
			updateTaskHandler(w, r)
		case http.MethodDelete:
			deleteTaskHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
