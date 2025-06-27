package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"GoTST/internal/database"
	"GoTST/internal/handlers"
	"GoTST/internal/tasksService"
	"GoTST/internal/userService"
	"GoTST/internal/web/tasks"
	"GoTST/internal/web/users"
)

func main() {
	// Инициализируем БД
	database.InitDB()
	if err := database.DB.AutoMigrate(&tasksService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	// Сервисы и хендлеры для задач
	tasksRepo := tasksService.NewTaskRepository(database.DB)
	tasksService := tasksService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	// Сервисы и хендлеры для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем обработчики
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlersWithBaseURL(e, tasksStrictHandler, "")

	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlersWithBaseURL(e, usersStrictHandler, "")

	// Запуск
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start: %v", err)
	}
}
