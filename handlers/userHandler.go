package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func GetUserHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	rows, err := db.Query(ctx, `SELECT id,name,email,created_at,updated_at FROM "Users" WHERE id = $1`, userId)
	defer rows.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}

	user, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UserResponse])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to collect user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}
