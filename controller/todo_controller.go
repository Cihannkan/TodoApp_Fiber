package controller

import (
	"TodoApp/model"   // Modül yolunu güncelleyin
	"TodoApp/service" // Modül yolunu güncelleyin
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	Service *service.TodoService
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo model.Todo
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if todo.Task == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Task is required",
		})
	}
	fmt.Println(todo)
	if err := c.Service.CreateTodo(&todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"id":      todo.ID,
		"task":    todo.Task,
		"done":    todo.Done,
		"message": "Todo created successfully",
	})
}
