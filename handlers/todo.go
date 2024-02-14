package handlers

import (
	"makromusic-task/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTodoListHandler(c *fiber.Ctx) error {
	page := c.Query("page", "1")
	pageSize := c.Query("pageSize", "10")

	pageNum, _ := strconv.Atoi(page)
	size, _ := strconv.Atoi(pageSize)

	todos, err := models.GetTodoList(pageNum, size)
	if err != nil {
		return err
	}

	return c.JSON(todos)
}
func AddTodoHandler(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	err := todo.Create()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Todo created successfully"})
}

func CompleteTodoHandler(c *fiber.Ctx) error {
	todoID := c.Params("id")
	err := models.MarkTodoAsComplete(todoID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Todo marked as complete"})
}
