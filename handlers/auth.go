package handlers

import (
	"makromusic-task/models"
	"makromusic-task/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := user.Create()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func LoginHandler(c *fiber.Ctx) error {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&credentials); err != nil {
		return err
	}

	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if err := utils.ComparePasswords(user.Password, credentials.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}
