package main

import (
	"appauths/initializers"
	"appauths/routes/appRoutes"
	"appauths/routes/authRoutes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	if err := initializers.InitDBPool(); err != nil {
		log.Fatalln(err)
	}

	initializers.InitSessionStores()

	initializers.InitOauth2Config()
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Route("/api/auth/signup", authRoutes.Signup)

	app.Route("/api/auth/login", authRoutes.Login)

	app.Route("/api/auth/oauth", authRoutes.OAuth)

	app.Route("/api/app", appRoutes.App)

	app.Listen(":5000")
}
