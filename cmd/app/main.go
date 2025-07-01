package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/SaidGo/Test-golang/internal/database"
	"github.com/SaidGo/Test-golang/internal/handlers"
	"github.com/SaidGo/Test-golang/internal/tasksService"
	"github.com/SaidGo/Test-golang/internal/userService"
	"github.com/SaidGo/Test-golang/internal/web/tasks"
	"github.com/SaidGo/Test-golang/internal/web/users"
)

func main() {
	// Инициализируем БД
	database.InitDB()
	if err := database.DB.AutoMigrate(&tasksService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	// Сервисы и хендлеры для задач
	tasksRepo := tasksService.NewTaskRepository(database.DB)
	var tasksSvc tasksService.Service = tasksService.NewService(tasksRepo) // интерфейс
	tasksHandler := handlers.NewHandler(tasksSvc)

	// Сервисы и хендлеры для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	var userSvc userService.Service = userService.NewService(userRepo) // интерфейс
	userHandler := handlers.NewUserHandler(userSvc)

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем обработчики без NewStrictHandler
	tasks.RegisterHandlersWithBaseURL(e, tasksHandler, "")
	users.RegisterHandlersWithBaseURL(e, userHandler, "")

	// Запуск
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
