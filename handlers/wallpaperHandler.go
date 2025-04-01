package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func WallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx := context.Background()
	db := utils.Database()
	defer db.Close(ctx)

	query := `
		SELECT w.*, a.name AS artist_name, c.name AS category_name 
		FROM "Wallpapers" w 
		JOIN "Artists" a ON w."artistsId" = a.id 
		JOIN "Categories" c ON w."categoriesId" = c.id
	`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve wallpapers from the database",
		})
	}
	defer rows.Close()

	wallpapers, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Wallpapers])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process wallpapers data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"wallpapers": wallpapers,
	})
}

func ParticularWallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx := context.Background()
	db := utils.Database()
	defer db.Close(ctx)

	wallpaperId := c.Params("id")
	if wallpaperId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wallpaper ID is required",
		})
	}

	row, err := db.Query(ctx, `SELECT w.*, a.name AS artist_name, c.name AS category_name FROM "Wallpapers" w JOIN "Artists" a ON w."artistsId" = a.id JOIN "Categories" c ON w."categoriesId" = c.id WHERE w.id = $1`, wallpaperId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve wallpaper",
		})
	}
	defer row.Close()

	wallpaper, err := pgx.CollectRows(row, pgx.RowToStructByName[models.Wallpapers])

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Wallpaper not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"wallpaper": wallpaper[0],
	})
}
