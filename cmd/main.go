package main

import (
	"fmt"
	"log"
	"project_x/internal/db"
	"project_x/internal/handlers"
	"project_x/internal/taskService"
	"project_x/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("could not connect to DB: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskServ := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewHandler(taskServ)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	fmt.Println("Server started at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
