package authRoutes

import (
	"appauths/handlers/signupHandlers"

	"github.com/gofiber/fiber/v2"
)

func Signup(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the signup process
	router.Post("/request_new_account", signupHandlers.RequestNewAccount) // issue cookie associated with the next state info
	router.Post("/verify_email", signupHandlers.VerifyEmail)              // verify cookie and state, update cookie for next state
	router.Post("/register_user", signupHandlers.RegisterUser)            // verify cookie and state, destroy cookie
}
