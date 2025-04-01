package main

import (
	_ "Pixelbloom-Backend/docs"
	"Pixelbloom-Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
)

// @title 						Pixelbloom Appliaction API
// @version 					1.0
// @description 				This is the API documentation for the Pixelbloom application.
// @host 						localhost:3000
// @BasePath 					/api
// @schemes 					http
// @securityDefinitions.apikey 	Bearer
// @in 							header
// @name 						Authorization
// @contact.name 				Pixelbloom Team
// @contact.url					https://GitHub.com/Jenil-Desai/Pixelbloom-Backend
// @contact.email				jenildev91@gmail.com
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	api := app.Group("/api")

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:          "http://localhost:3000/doc.json",
		DeepLinking:  false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	api.Route("/auth", routes.AuthRouter)
	api.Route("/wallpapers", routes.WallpaperRouter)
	api.Route("/wallpapers", routes.LikeRouter)

	log.Fatal(app.Listen(":3000"))
}
