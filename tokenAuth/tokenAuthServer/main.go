package main

import (
	"codeauths/helpers"
	"codeauths/tokenAuth/tokenAuthServer/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	if err := helpers.LoadEnv(".env"); err != nil {
		log.Fatalln(err)
	}

	if err := helpers.InitDBPool(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// signup (verify email) : OTP
	app.Route("/api/auth/signup", handlers.Signup)

	// login

	// access a restricted resource : jwt authentication

	// change your password : OTP

	// make payment : live token

	app.Listen(":5000")
}
