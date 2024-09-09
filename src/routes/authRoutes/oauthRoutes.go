package authRoutes

import (
	"appauths/src/handlers/oauthHandlers"

	"github.com/gofiber/fiber/v2"
)

func GoogleOAuth(router fiber.Router) {
	router.Get("/auth_url", oauthHandlers.GoogleAuthURL)

	router.Get("/callback", oauthHandlers.GoogleAuthCallback)
}

func GithubOAuth(router fiber.Router) {
	router.Get("/auth_url", oauthHandlers.GithubAuthURL)
	router.Get("/callback", oauthHandlers.GithubAuthCallback)
}

func OAuth(router fiber.Router) {

	router.Route("/google", GoogleOAuth)
	router.Route("/github", GithubOAuth)
}
