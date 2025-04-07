package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// @summary		Get Wallpapers
// @description	Retrieve all wallpapers
// @tags		wallpapers
// @accept		json
// @produce		json
// @success		200	{object}	map[string]interface{}	"List of wallpapers"
// @failure		401	{object}	map[string]interface{}	"Unauthorized"
// @failure		500	{object}	map[string]interface{}	"Failed to retrieve wallpapers from the database"
// @router		/wallpapers [get]
// @security		Bearer
func WallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx := context.Background()
	db := utils.Database()
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			fmt.Printf("Failed to close database connection %v", err)
		}
	}(db, ctx)

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

	return c.Status(fiber.StatusOK).JSON(wallpapers)
}

// @summary		Get Wallpaper by ID
// @description	Retrieve a specific wallpaper by its ID
// @tags		wallpapers
// @accept		json
// @produce		json
// @param		id	path	string	true	"Wallpaper ID"
// @success		200	{object}	map[string]interface{}	"Wallpaper details"
// @failure		400	{object}	map[string]interface{}	"Wallpaper ID is required"
// @failure		401	{object}	map[string]interface{}	"Unauthorized"
// @failure		404	{object}	map[string]interface{}	"Wallpaper not found"
// @failure		500	{object}	map[string]interface{}	"Failed to retrieve wallpaper"
// @router		/wallpapers/{id} [get]
// @Security 	Bearer
func ParticularWallpaperHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx := context.Background()
	db := utils.Database()
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			fmt.Printf("Failed to close database connection %v", err)
		}
	}(db, ctx)

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

	if err != nil || len(wallpaper) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Wallpaper not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(wallpaper[0])
}
