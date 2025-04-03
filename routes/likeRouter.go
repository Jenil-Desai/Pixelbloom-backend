package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func LikeRouter(api fiber.Router) {
	api.Get("/liked-wallpapers", middlewares.VerifyToken, handlers.GetLikedWallpapersHandler)
	api.Post("/:id/like", middlewares.VerifyToken, handlers.LikeWallpaperHandler)
	api.Delete("/:id/like", middlewares.VerifyToken, handlers.UnlikeWallpaperHandler)
}
