package signupHandlers

import (
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("step").(string) != "signup: register user" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	// check if user with username already exists
	userExists, err := helpers.QueryRowField[bool]("SELECT EXISTS(SELECT 1 FROM auth_user WHERE username = $1)", body.Username)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if *userExists {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("username already taken")
	}

	res, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	hashPwd := string(res)

	_, dbin_err := helpers.QueryRowField[bool]("INSERT INTO auth_user (email, username, password) VALUES ($1, $2, $3) RETURNING true", session.Get("email").(string), body.Username, hashPwd)
	if dbin_err != nil {
		log.Println(dbin_err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if sd_err := session.Destroy(); sd_err != nil {
		panic(sd_err)
	}

	return c.SendString("Registration success. Proceed to login.\n")
}
