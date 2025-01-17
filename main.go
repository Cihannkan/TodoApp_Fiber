package main

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func main() {
	app := fiber.New()
	app.Post("/user", func(c *fiber.Ctx) error {
		var Todo Todo
		if err := c.BodyParser(&Todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}
		return c.JSON(Todo)
	})

	app.Listen(":3000")
}
