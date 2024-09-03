package authRoutes

import "github.com/gofiber/fiber/v2"

func GoogleOAuth(router fiber.Router) {
	// generate and authorization URL
	router.Get("/auth_url", nil)

	// exchange authorization code for tokens (access_token & refresh_token)
	// save both tokens (hashed) in the database
	router.Get("/callback", nil)

	// when needed, use refresh_token to get a new access_token
	router.Post("/refresh_token", nil)

	// use access token to request for user profile info
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
