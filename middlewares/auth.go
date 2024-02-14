package middlewares

import (
	"makromusic-task/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return fiber.ErrUnauthorized
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	c.Locals("userID", userID)

	return c.Next()
}
