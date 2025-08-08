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

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Text   string `json:"task"`
	IsDone bool   `gorm:"default:false" json:"default:false"`
}

func initDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=tasks port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database", err)
	}

	db.AutoMigrate(&Task{})
}

func getTasks(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func createTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid request"})
	}

	task.ID = uuid.NewString()

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")
	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "falied to update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", createTask)
	e.POST("/tasks", getTasks)
	e.PATCH("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	fmt.Println("http://localhost:8080")
	e.Start(":8080")

}
