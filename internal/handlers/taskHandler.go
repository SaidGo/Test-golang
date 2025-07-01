package handlers

import (
	"net/http"

	"github.com/SaidGo/Test-golang/internal/tasksService"
	"github.com/SaidGo/Test-golang/internal/web/tasks"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service tasksService.Service
}

func NewHandler(s tasksService.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetTasks(ctx echo.Context) error {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]tasks.Task, 0, len(allTasks))
	for _, t := range allTasks {
		id := t.ID
		taskText := t.Text
		isDone := t.IsDone
		response = append(response, tasks.Task{
			Id:     &id,
			Task:   &taskText,
			IsDone: &isDone,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) PostTasks(ctx echo.Context) error {
	var reqBody tasks.PostTasksJSONRequestBody
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if reqBody.Task == nil || reqBody.IsDone == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "missing task or is_done"})
	}

	task := tasksService.Task{
		Text:   *reqBody.Task,
		IsDone: *reqBody.IsDone,
	}

	created, err := h.service.CreateTask(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := tasks.Task{
		Id:     &created.ID,
		Task:   &created.Text,
		IsDone: &created.IsDone,
	}

	return ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) PatchTasksId(ctx echo.Context, id int) error {
	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	var reqBody tasks.PatchTasksIdJSONRequestBody
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if reqBody.Task != nil {
		task.Text = *reqBody.Task
	}
	if reqBody.IsDone != nil {
		task.IsDone = *reqBody.IsDone
	}

	updated, err := h.service.UpdateTask(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	resp := tasks.Task{
		Id:     &updated.ID,
		Task:   &updated.Text,
		IsDone: &updated.IsDone,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteTasksId(ctx echo.Context, id int) error {
	err := h.service.DeleteTask(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
