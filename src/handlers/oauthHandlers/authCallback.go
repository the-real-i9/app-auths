package oauthHandlers

import (
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

func GoogleAuthCallback(c *fiber.Ctx) error {
	ctx := context.Background()

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	state, err := helpers.JwtVerify[string](c.Query("state"), os.Getenv("SESSION_JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	if session.Get("state").(string) != *state {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	authCode := c.Query("code")

	verifier := session.Get("verifier").(string)

	token, err := globalVars.Oauth2Config.Exchange(ctx, authCode, oauth2.VerifierOption(verifier))
	if err != nil {
		panic(err)
	}

	service, err := people.NewService(ctx, option.WithTokenSource(globalVars.Oauth2Config.TokenSource(ctx, token)))
	if err != nil {
		panic(err)
	}

	person, err := service.People.Get("people/me").Fields("names", "nicknames", "emailAddresses").Do()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// check if the user already has an account,
	// if yes, return a JWT token (log the user in)
	// if no, sign the user in with a random username (user can change it later)

	return c.JSON(person)
}
