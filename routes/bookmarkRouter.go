package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func BookmarkRouter(api fiber.Router) {
	api.Get("/bookmarks", middlewares.VerifyToken, handlers.GetBookmarks)
	api.Post("/:id/bookmark", middlewares.VerifyToken, handlers.BookmarkWallpaperHandler)
	api.Delete("/:id/bookmark", middlewares.VerifyToken, handlers.UnbookmarkWallpaperHandler)
}
