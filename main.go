package main

import (
	"Pixelbloom-Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	api := app.Group("/api")

	api.Route("/auth", routes.AuthRouter)
	api.Route("/wallpapers", routes.WallpaperRouter)
	api.Route("/wallpapers", routes.LikeRouter)

	log.Fatal(app.Listen(":3000"))
}
