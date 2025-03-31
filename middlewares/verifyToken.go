package middlewares

import (
	"Pixelbloom-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	// Split the header to get the token part
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := parts[1]

	if res, err := utils.VerifyToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized: " + err.Error(),
		})
	} else {
		c.Locals("userId", res["user_id"])

		return c.Next()
	}
}
