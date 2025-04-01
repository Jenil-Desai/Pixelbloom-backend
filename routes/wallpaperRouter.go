package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func WallpaperRouter(api fiber.Router) {
	api.Get("/", middlewares.VerifyToken, handlers.WallpaperHandler)
	api.Get("/:id", middlewares.VerifyToken, handlers.ParticularWallpaperHandler)
}
