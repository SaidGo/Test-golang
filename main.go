package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var task string // глобальная переменная

type TaskRequest struct {
	Task string `json:"task"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка JSON", http.StatusBadRequest)
		return
	}
	task = req.Task
	fmt.Fprintf(w, "Задача сохранена: %s", task)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", task)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postHandler(w, r)
		case http.MethodGet:
			getHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
