package main

import (
	"appauths/globalVars"
	"appauths/helpers"
	"appauths/routes"
	"appauths/routes/authRoutes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	if err := helpers.InitDBPool(); err != nil {
		log.Fatalln(err)
	}

	storage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_auth"})

	globalVars.AuthSessionStore = session.New(session.Config{
		Storage:    storage,
		CookiePath: "/api/auth",
	})
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Route("/api/auth/signup", authRoutes.Signup)

	app.Route("/api/auth/login", authRoutes.Login)

	app.Route("/api/auth/oauth", authRoutes.OAuth)

	app.Route("/api/app", routes.App)

	app.Listen(":5000")
}
