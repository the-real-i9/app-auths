package signupHandlers

import (
	"codeauths/helpers"
	"codeauths/tokenAuth/tokenAuthServer/globalVars"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SubmitEmail(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	// check if user with email already exists
	userExists, err := helpers.QueryRowField[bool]("SELECT EXISTS(SELECT 1 FROM tokenauth.i9ca_user WHERE email = $1)", body.Email)
	if err != nil {
		log.Error(err)
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

	session.Set("email", body.Email)
	session.Set("verificationToken", verfToken)
	session.Set("verificationTokenExpires", verfTokenExpires)

	if save_err := session.Save(); save_err != nil {
		panic(save_err)
	}

	go helpers.SendMail(body.Email, "Verify your email", fmt.Sprintln("Your email verification code is", verfToken))

	return c.SendString("Email verification code has been sent to " + body.Email)
}

func VerifyEmail(c *fiber.Ctx) error {
	var body struct {
		InputVerfToken int `json:"verification_code"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	session, err := globalVars.SignupSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	email := session.Get("email").(string)
	verfToken := session.Get("verificationToken").(int)
	verfTokenExpires := time.Unix(session.Get("verificationTokenExpires").(int64), 0)

	if verfToken != body.InputVerfToken {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Incorrect verification code")
	}

	if time.Now().After(verfTokenExpires) {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Verification code expired")
	}

	return c.SendString("Your email " + email + " has been verified!")
}

func RegisterUser(c *fiber.Ctx) error {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	// check if user with username already exists
	userExists, err := helpers.QueryRowField[bool]("SELECT EXISTS(SELECT 1 FROM tokenauth.i9ca_user WHERE username = $1)", body.Username)
	if err != nil {
		log.Error(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if *userExists {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("username already taken")
	}

	session, err := globalVars.SignupSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	email := session.Get("email").(string)

	_, dbin_err := helpers.QueryRowField[bool]("INSERT INTO tokenauth.i9ca_user (email, username, password) VALUES ($1, $2, $3)", email, body.Username, body.Password)
	if dbin_err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if sd_err := session.Destroy(); sd_err != nil {
		panic(sd_err)
	}

	return c.SendString("Registration success. Proceed to login.")
}
