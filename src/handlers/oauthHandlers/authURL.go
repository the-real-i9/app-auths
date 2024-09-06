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
	// return c.Redirect(url)
}
