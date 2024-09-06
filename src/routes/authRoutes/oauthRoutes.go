package authRoutes

import (
	"appauths/src/handlers/oauthHandlers"

	"github.com/gofiber/fiber/v2"
)

func GoogleOAuth(router fiber.Router) {
	router.Get("/auth_url", oauthHandlers.GoogleAuthURL)

	router.Get("/callback", oauthHandlers.GoogleAuthCallback)

	// use access token to request for user profile info
	// if access token has expired, use refresh_token to get a new access_token
	router.Get("/", nil)
}

func GithubOAuth(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Get("/callback", nil)
}

func OAuth(router fiber.Router) {

	router.Route("/google", GoogleOAuth)
	router.Route("/github", GithubOAuth)
}
