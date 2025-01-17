package main

import (
	"TodoApp/controller"
	"TodoApp/model"
	"TodoApp/repository"
	"TodoApp/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Todo{})

	todoRepo := &repository.TodoRepository{DB: db}
	todoService := &service.TodoService{Repo: todoRepo}
	todoController := &controller.TodoController{Service: todoService}

	app.Post("/user", todoController.CreateTodo)

	app.Listen(":3000")
}
