package handlers

import (
	"net/http"

	"github.com/SaidGo/Test-golang/internal/userService"
	"github.com/SaidGo/Test-golang/internal/web/users"
	"github.com/SaidGo/Test-golang/internal/web/tasks"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service userService.Service
}

func NewUserHandler(service userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	usersList, err := h.service.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]users.User, 0, len(usersList))
	for _, u := range usersList {
		id := int(u.ID)
		response = append(response, users.User{
			Id:       &id,
			Email:    &u.Email,
			Password: &u.Password,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) PostUser(ctx echo.Context) error {
	var input users.PostUserJSONRequestBody
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	user := &userService.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.CreateUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	id := int(user.ID)
	return ctx.JSON(http.StatusCreated, users.User{
		Id:       &id,
		Email:    &user.Email,
		Password: &user.Password,
	})
}

func (h *UserHandler) PatchUserByID(ctx echo.Context, id int) error {
	var input users.PatchUserByIDJSONRequestBody
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	user := &userService.User{
		ID:       uint(id),
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.service.UpdateUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	respID := int(user.ID)
	return ctx.JSON(http.StatusOK, users.User{
		Id:       &respID,
		Email:    &user.Email,
		Password: &user.Password,
	})
}

func (h *UserHandler) DeleteUserByID(ctx echo.Context, id int) error {
	if err := h.service.DeleteUser(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

// GET /users/{id}/tasks
func (h *UserHandler) GetTasksByUserID(ctx echo.Context, id int) error {
	tasksList, err := h.service.GetTasksForUser(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]tasks.Task, 0, len(tasksList))
	for _, t := range tasksList {
		taskID := t.ID
		text := t.Text
		done := t.IsDone
		userID := t.UserID
		response = append(response, tasks.Task{
			Id:     &taskID,
			Task:   &text,
			IsDone: &done,
			UserId: &userID,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
