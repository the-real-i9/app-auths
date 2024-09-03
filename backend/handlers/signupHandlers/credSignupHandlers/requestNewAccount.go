package credSignupHandlers

import (
	"appauths/backend/globalVars"
	"appauths/backend/helpers"
	"fmt"
	"math/rand"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
)

func RequestNewAccount(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(fiber.ErrUnprocessableEntity.Message)
	}

	// check if user with email already exists
	userExists, err := helpers.QueryRowField[bool]("SELECT EXISTS(SELECT 1 FROM auth_user WHERE email = $1)", body.Email)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if *userExists {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("User with email already exists")
	}

	verfToken := rand.Intn(899999) + 100000
	verfTokenExpires := time.Now().Add(30 * time.Minute).Unix()

	// create signup session with verification token
	session, err := globalVars.SignupSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	go helpers.SendMail(body.Email, "Verify your email", fmt.Sprintln("Your email verification code is", verfToken))

	session.Set("email", body.Email)
	session.Set("verificationToken", verfToken)
	session.Set("verificationTokenExpires", verfTokenExpires)
	session.Set("step", "verify email")

	if save_err := session.Save(); save_err != nil {
		panic(save_err)
	}

	return c.SendString("Email verification code has been sent to " + body.Email + "\n")
}
