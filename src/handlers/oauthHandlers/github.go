package oauthHandlers

import (
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func GithubAuthURL(c *fiber.Ctx) error {
	verifier := oauth2.GenerateVerifier()

	state := helpers.JwtSign("oauth: github callback", os.Getenv("SESSION_JWT_SECRET"), time.Now().Add(24*time.Hour))

	url := globalVars.GithubOauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	session.Set("state", "oauth: github callback")
	session.Set("verifier", verifier)

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendString(fmt.Sprintf("Visit the URL for the auth dialog: %v", url))
}

func GithubAuthCallback(c *fiber.Ctx) error {
	ctx := context.Background()

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	state, err := helpers.JwtVerify[string](c.Query("state"), os.Getenv("SESSION_JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	sessState, _ := session.Get("state").(string)
	if sessState != *state {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	authCode := c.Query("code")

	verifier := session.Get("verifier").(string)

	token, err := globalVars.GithubOauth2Config.Exchange(ctx, authCode, oauth2.VerifierOption(verifier))
	if err != nil {
		panic(err)
	}

	// token.Accesstoken to make request
	agent := fiber.Get("https://api.github.com/user").Set("Authorization", "Bearer "+token.AccessToken)
	_, data, errs := agent.Bytes()
	if len(errs) > 0 {
		log.Println(errs)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(string(data))

	/* userEmail := ""
	// check if the user already has an account,
	user, err := helpers.QueryRowType[appTypes.User]("SELECT user_id, email, username FROM auth_user WHERE email = $1", userEmail)
	if err != nil {
		log.Println(err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	// if yes, return a JWT token (log the user in)
	if user != nil {
		jwt := helpers.JwtSign(user, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

		return c.JSON(fiber.Map{
			"msg":     "Login success!",
			"authJwt": jwt,
		})
	}

	// if no, sign up the user using the portion before the @ on the user's email
	// (invalid characters replaced with "_") (user can change it later)
	strRep := strings.NewReplacer(".", "_", "-", "_")
	tempUsername := strRep.Replace(strings.Split(userEmail, "@")[0])
	newUser, err := helpers.QueryRowType[appTypes.User]("INSERT INTO auth_user (email, username) VALUES ($1, $2) RETURNING user_id, email, username", userEmail, tempUsername)
	if err != nil {
		log.Println(err)
		c.SendStatus(fiber.StatusInternalServerError)
	}

	// return a JWT token (log the user in)
	jwt := helpers.JwtSign(newUser, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

	if err := session.Destroy(); err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"msg":     "Login success!",
		"authJwt": jwt,
	}) */
}
