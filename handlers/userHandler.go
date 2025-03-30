package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(c *fiber.Ctx) error {
	body := new(SignUpRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	db := utils.Database()

	hashedPassword, err := utils.HashPassword(body.Password)

	id := uuid.NewString()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	db.Create(&models.User{ID: id, Email: body.Email, Password: hashedPassword})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignInHandler(c *fiber.Ctx) error {
	body := new(SignInRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	return nil
}
