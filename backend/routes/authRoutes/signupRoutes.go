package authRoutes

import (
	"appauths/backend/handlers/signupHandlers/credSignupHandlers"

	"github.com/gofiber/fiber/v2"
)

func CredSignup(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the signup process
	router.Post("/request_new_account", credSignupHandlers.RequestNewAccount) // issue cookie associated with the next state info
	router.Post("/verify_email", credSignupHandlers.VerifyEmail)              // verify cookie and state, update cookie for next state
	router.Post("/register_user", credSignupHandlers.RegisterUser)            // verify cookie and state, destroy cookie
}

func GoogleSignup(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Post("/", nil)
}

func GithubSignup(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Post("/", nil)
}

func Signup(router fiber.Router) {
	router.Route("/with_cred", CredSignup)

	router.Route("/oauth/with_google", GoogleSignup)
	router.Route("/oauth/with_github", GithubSignup)
}
