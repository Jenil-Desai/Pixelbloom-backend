package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func GetBookmarks(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	rows, err := db.Query(ctx, `SELECT w.*, a.name AS artist_name, c.name AS category_name FROM "Wallpapers" w JOIN "BookmarkedWallpapers" bw ON w.id = bw."wallpaperId" JOIN "Artists" a ON w."artistsId" = a.id JOIN "Categories" c ON w."categoriesId" = c.id WHERE bw."userId" = $1`, userId)
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

	return c.Status(200).JSON(wallpapers)
}

func BookmarkWallpaperHandler(c *fiber.Ctx) error {
	wallpaperId := c.Params("id")
	userId := c.Locals("userId").(string)
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	_, err := db.Exec(ctx, `INSERT INTO "BookmarkedWallpapers" ("userId", "wallpaperId") VALUES ($1, $2)`, userId, wallpaperId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Wallpaper bookmarked successfully",
	})
}

func UnbookmarkWallpaperHandler(c *fiber.Ctx) error {
	wallpaperId := c.Params("id")
	userId := c.Locals("userId").(string)
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	_, err := db.Exec(ctx, `DELETE FROM "BookmarkedWallpapers" WHERE "userId" = $1 AND "wallpaperId" = $2`, userId, wallpaperId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Wallpaper unbookmarked successfully",
	})
}
