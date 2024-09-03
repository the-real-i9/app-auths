package authRoutes

import "github.com/gofiber/fiber/v2"

func GoogleOAuth(router fiber.Router) {
	// generate authorization URL
	router.Get("/auth_url", nil)

	// exchange authorization code for access token
	router.Get("/callback", nil)

	// when needed, refresh token for new access token
	router.Get("/refresh_token", nil)

	// use access token to request for user profile info
	router.Get("/", nil)
}

func GithubOAuth(router fiber.Router) {
	router.Get("/auth_url", nil)
	router.Get("/callback", nil)
}

func OAuth(router fiber.Router) {

	router.Route("/oauth/google", GoogleOAuth)
	router.Route("/oauth/github", GithubOAuth)
}
