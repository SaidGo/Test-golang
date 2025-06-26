package handlers

import (
	"context"

	"GoTST/internal/tasksService"
	"GoTST/internal/web/tasks"
)

type Handler struct {
	service *tasksService.Service
}

func NewHandler(s *tasksService.Service) *Handler {
	return &Handler{service: s}
}

// GET /tasks
func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := make(tasks.GetTasks200JSONResponse, 0)
	for _, t := range allTasks {
		task := tasks.Task{
			Id:     &t.ID,
			Task:   &t.Text,
			IsDone: &t.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// POST /tasks
func (h *Handler) PostTasks(_ context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskReq := req.Body

	task := tasksService.Task{
		Text:   *taskReq.Task,
		IsDone: *taskReq.IsDone,
	}

	created, err := h.service.CreateTask(task)
	if err != nil {
		return nil, err
	}

	resp := tasks.PostTasks201JSONResponse{
		Id:     &created.ID,
		Task:   &created.Text,
		IsDone: &created.IsDone,
	}

	return resp, nil
}

// PATCH /tasks/{id}
func (h *Handler) PatchTasksId(_ context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := req.Id

	task, err := h.service.GetTaskByID(uint(id))
	if err != nil {
		return nil, err
	}

	body := req.Body

	if body.Task != nil {
		task.Text = *body.Task
	}
	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	}

	updated, err := h.service.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	resp := tasks.PatchTasksId200JSONResponse{
		Id:     &updated.ID,
		Task:   &updated.Text,
		IsDone: &updated.IsDone,
	}

	return resp, nil
}

// DELETE /tasks/{id}
func (h *Handler) DeleteTasksId(_ context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.service.DeleteTask(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
