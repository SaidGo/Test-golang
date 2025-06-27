package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"GoTST/internal/database"
	"GoTST/internal/handlers"
	"GoTST/internal/tasksService"
	"GoTST/internal/web/tasks"

)

func main() {
	// Инициализируем БД
	database.InitDB()
	if err := database.DB.AutoMigrate(&tasksService.Task{}); err != nil {
	log.Fatalf("auto-migrate failed: %v", err)
}


	// Сервисы и хендлер
	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем StrictHandler
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlersWithBaseURL(e, strictHandler, "")


	// Запуск
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
