package main

import (
	"fmt"
	"log"
	"project_x/internal/db"
	"project_x/internal/handlers"
	"project_x/internal/taskService"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}

	if err := database.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("could not migrate: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskServ := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskServ)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandler.GetTasks)
	e.POST("/tasks", taskHandler.PostTasks)
	e.PATCH("/tasks/:id", taskHandler.PatchTasks)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	fmt.Println("Server started at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
