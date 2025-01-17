package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Task string `json:"task"`
	Done bool   `json:"done" gorm:"default:false"`
}

func main() {
	app := fiber.New()
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})

	app.Post("/user", func(c *fiber.Ctx) error {
		var todo Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}
		if todo.Task == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Task is required",
			})
		}
		fmt.Println(todo)
		if result := db.Create(&todo); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"id":      todo.ID,
			"task":    todo.Task,
			"done":    todo.Done,
			"message": "Todo created successfully",
		})
	})

	app.Listen(":3000")
}
