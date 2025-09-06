package handlers

import (
	"net/http"
	"project_x/internal/taskService"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PostTasks(c echo.Context) error {
	var req struct {
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task, err := h.service.CreateTask(req.Text, req.IsDone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchTasks(c echo.Context) error {
	id := c.Param("id")

	var req struct {
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedTask, err := h.service.UpdateTask(id, req.Text, req.IsDone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
