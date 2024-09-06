package oauthHandlers

import (
	"appauths/src/appTypes"
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"context"
	"log"
	"os"
	"strings"
	"time"

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

	sessState, _ := session.Get("state").(string)
	if sessState != *state {
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

	person, err := service.People.Get("people/me").PersonFields("emailAddresses").Do()
	if err != nil {
		panic(err)
	}

	userEmail := person.EmailAddresses[0].Value
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
	})
}
