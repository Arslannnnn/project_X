package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=tasks port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("could not migrate: %v", err)
	}
}

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Text   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func postTask(c echo.Context) error {
	var req struct {
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	task := Task{
		ID:     uuid.NewString(),
		Text:   req.Text,
		IsDone: req.IsDone,
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not add task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func getTasks(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func putTask(c echo.Context) error {
	id := c.Param("id")

	var req struct {
		Text   string `json:"task"`
		IsDone bool   `json:"is_done"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "could not find task"})
	}

	task.Text = req.Text
	task.IsDone = req.IsDone

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTask)
	e.PATCH("/tasks/:id", putTask)
	e.DELETE("/tasks/:id", deleteTask)

	fmt.Println("http://localhost:8080")
	e.Start(":8080")
}
