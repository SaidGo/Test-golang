package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID   int    `json:"id"`   
	Text string `json:"text"` 
	Done bool   `json:"done"` 
}

var tasks = make(map[int]Task) 
var nextID = 1                 


func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	task := Task{
		ID:   nextID,
		Text: input.Text,
		Done: false,
	}
	tasks[nextID] = task
	nextID++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}


func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	var list []Task
	for _, t := range tasks {
		list = append(list, t)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}


func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	var input struct {
		Text *string `json:"text,omitempty"`
		Done *bool   `json:"done,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if input.Text != nil {
		task.Text = *input.Text
	}
	if input.Done != nil {
		task.Done = *input.Done
	}

	tasks[id] = task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}


func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	if _, exists := tasks[id]; !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	delete(tasks, id)
	w.WriteHeader(http.StatusNoContent)
}


func main() {
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTaskHandler(w, r)
		case http.MethodGet:
			listTasksHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
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
