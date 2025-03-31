package routes

import (
	"Pixelbloom-Backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Post("/signup", handlers.SignUpHandler)
	api.Post("/signin", handlers.SignInHandler)
}
