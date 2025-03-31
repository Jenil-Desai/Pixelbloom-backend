package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Post("/signup", handlers.SignUpHandler)
	api.Post("/signin", handlers.SignInHandler)

	api.Get("/likedWallpapers", middlewares.VerifyToken, handlers.GetLikedWallpapersHandler)
}
