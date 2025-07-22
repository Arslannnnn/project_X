package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   string `json:"id"`
	Text string `json:"task"`
}

var tasks []Task

func getTask(c echo.Context) error {
	if tasks == nil {
		return c.JSON(http.StatusOK, []Task{})
	}
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var request struct {
		Task string `json: "task"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := Task{
		ID:   uuid.NewString(),
		Text: request.Task,
	}
	tasks = append(tasks, task)

	return c.JSON(http.StatusCreated, task)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var request struct {
		Task string `json: "task"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Text = request.Task
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	fmt.Println("сервер запущен")

	e.Logger.Fatal(e.Start("localhost:8080"))
}
