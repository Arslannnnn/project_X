package handlers

import (
	"context"
	"project_x/internal/taskService"
	"project_x/internal/web/tasks"
)

type Handler struct {
	service taskService.TaskService
}

func NewHandler(s taskService.TaskService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     tsk.ID,
			Task:   tsk.Text,
			IsDone: tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	createdTask, err := h.service.CreateTask(taskRequest.Task, taskRequest.IsDone)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     createdTask.ID,
		Task:   createdTask.Text,
		IsDone: createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.service.DeleteTask(request.Id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskRequest := request.Body

	updatedTask, err := h.service.UpdateTask(request.Id, taskRequest.Task, taskRequest.IsDone)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     updatedTask.ID,
		Task:   updatedTask.Text,
		IsDone: updatedTask.IsDone,
	}
	return response, nil
}
