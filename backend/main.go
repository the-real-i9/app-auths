package main

import (
	"i9codesauths/backend/globalVars"
	"i9codesauths/backend/handlers"
	"i9codesauths/backend/helpers"
	"i9codesauths/backend/routes"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
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

	storage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_signup"})

	globalVars.SignupSessionStore = session.New(session.Config{
		Storage:    storage,
		CookiePath: "/api/auth/signup",
	})
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// signup | verify email : session auth | Email OTP auth
	app.Route("/api/auth/signup", routes.Signup)

	// password reset : session auth | OTP generator server

	// login : 2FA with TOTP | issue jwt
	app.Post("/api/auth/login", handlers.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("AUTH_JWT_SECRET"))},
	}))

	// access a restricted resource : jwt auth

	app.Listen(":5000")
}
