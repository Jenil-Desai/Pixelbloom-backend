package routes

import (
	"Pixelbloom-Backend/handlers"
	"Pixelbloom-Backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	api.Get("/", middlewares.VerifyToken, handlers.GetUserHandler)
}
