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

	authStorage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_auth"})
	appStorage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_process"})

	globalVars.AuthSessionStore = session.New(session.Config{
		Storage:    authStorage,
		CookiePath: "/api/auth",
	})

	globalVars.AppSessionStore = session.New(session.Config{
		Storage:    appStorage,
		CookiePath: "/api/app",
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
