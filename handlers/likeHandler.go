package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func GetLikedWallpapersHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	rows, err := db.Query(ctx, `SELECT w.* FROM "Wallpapers" w JOIN "LikedWallpapers" lw ON w.id = lw."wallpaperId" WHERE lw."userId" = $1`, userId)
	defer rows.Close()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}

	wallpapers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Wallpapers])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to collect wallpapers",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"wallpapers": wallpapers,
	})
}
