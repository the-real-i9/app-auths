package main

import (
	"codeauths/helpers"
	"codeauths/tokenAuth/tokenAuthServer/globalVars"
	"codeauths/tokenAuth/tokenAuthServer/handlers"
	"codeauths/tokenAuth/tokenAuthServer/routes"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
)

func init() {
	if err := helpers.LoadEnv("../../.env"); err != nil {
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

	// signup (verify email) : OTP
	app.Route("/api/auth/signup", routes.Signup)

	// change your password : OTP

	// login
	app.Post("/api/auth/login", handlers.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("AUTH_JWT_SECRET"))},
	}))

	// access a restricted resource : jwt authentication

	// make payment : live token

	app.Listen(":5000")
}
