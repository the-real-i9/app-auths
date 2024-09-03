package authRoutes

import (
	"appauths/backend/handlers/loginHandlers"

	"github.com/gofiber/fiber/v2"
)

func CredLogin(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the login process
	router.Post("/", loginHandlers.CredLogin)

}

func GoogleLogin(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Post("/", nil)
}

func GithubLogin(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Post("/", nil)
}

func MFALogin(router fiber.Router) {
	router.Post("/otp/email")

	router.Get("/totp/barcode", nil)
	router.Post("/totp/verify", nil)
}

func SSO(router fiber.Router) {

}

func Login(router fiber.Router) {
	router.Route("/with_cred", CredLogin)

	router.Route("/oauth/with_google", GoogleLogin)
	router.Route("/oauth/with_github", GithubLogin)

	router.Route("/mfa", MFALogin)

	router.Route("/sso", SSO)
}
