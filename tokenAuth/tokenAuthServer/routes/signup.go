package routes

import (
	"codeauths/tokenAuth/tokenAuthServer/handlers/signupHandlers"

	"github.com/gofiber/fiber/v2"
)

func Signup(router fiber.Router) {
	router.Post("/submit_email", signupHandlers.SubmitEmail)
	router.Post("/verify_email", signupHandlers.VerifyEmail)
	router.Post("/register_user", signupHandlers.RegisterUser)
}
