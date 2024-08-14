package authRoutes

import (
	"i9codesauths/backend/handlers/signupHandlers"

	"github.com/gofiber/fiber/v2"
)

func Signup(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the signup process
	router.Post("/submit_email", signupHandlers.SubmitEmail)   // issue cookie associated with the next state info
	router.Post("/verify_email", signupHandlers.VerifyEmail)   // verify cookie and state, update cookie for next state
	router.Post("/register_user", signupHandlers.RegisterUser) // verify cookie and state, destroy cookie
}
