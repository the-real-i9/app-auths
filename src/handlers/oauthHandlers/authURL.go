package oauthHandlers

import (
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func GoogleAuthURL(c *fiber.Ctx) error {
	verifier := oauth2.GenerateVerifier()

	state := helpers.JwtSign("oauth: google callback", os.Getenv("SESSION_JWT_SECRET"), time.Now().Add(24*time.Hour))

	url := globalVars.Oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	session.Set("state", "oauth: google callback")
	session.Set("verifier", verifier)

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendString(fmt.Sprintf("Visit the URL for the auth dialog: %v", url))
}

func GithubAuthURL(c *fiber.Ctx) error {

	clientId := os.Getenv("GITHUB_OAUTH_CLIENT_ID")
	redirectURI := "http://127.0.0.1:5000/api/auth/oauth/github/callback"
	scope := "user:email"
	state := helpers.JwtSign("oauth: github callback", os.Getenv("SESSION_JWT_SECRET"), time.Now().Add(24*time.Hour))

	url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=%s&state=%s", clientId, redirectURI, scope, state)

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	session.Set("state", "oauth: github callback")

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendString(fmt.Sprintf("Visit the URL for the auth dialog: %v", url))
}
