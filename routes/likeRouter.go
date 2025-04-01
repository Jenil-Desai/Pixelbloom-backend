package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func LikeRouter(api fiber.Router) {
	api.Get("/like-wallpaper", middlewares.VerifyToken, handlers.GetLikedWallpapersHandler)
}
