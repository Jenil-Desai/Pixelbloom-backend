package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"fmt"
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

	return c.Status(200).JSON(wallpapers)
}

func LikeWallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	wallpaperId := c.Params("id")
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `INSERT INTO "LikedWallpapers"("userId", "wallpaperId") VALUES($1, $2)`, userId, wallpaperId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert into LikedWallpapers",
		})
	}

	_, err = tx.Exec(ctx, `UPDATE "Wallpapers" SET "likes" = "likes" + 1 WHERE "id" = $1`, wallpaperId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update like count",
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Wallpaper liked successfully",
	})
}

func UnlikeWallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	wallpaperId := c.Params("id")
	ctx := context.Background()

	db := utils.Database()
	defer db.Close(ctx)

	tx, err := db.Begin(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM "LikedWallpapers" WHERE "userId" = $1 AND "wallpaperId" = $2`, userId, wallpaperId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete from LikedWallpapers",
		})
	}

	_, err = tx.Exec(ctx, `UPDATE "Wallpapers" SET "likes" = "likes" - 1 WHERE "id" = $1`, wallpaperId)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update like count",
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Wallpaper unliked successfully",
	})
}
